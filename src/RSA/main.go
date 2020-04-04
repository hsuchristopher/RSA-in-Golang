package main

import(
    "os"
    "fmt"
    "math/big"
    "math/rand"
    "time"
    "strconv"
)


/* Struct Declarations */
// This struct is used to hold my pub/priv keypair
type Keypair struct{
    key *big.Int
    n *big.Int
}

/*******************************************************************************
* This is where the program starts, Allows the User to enter two numbers then  *
* will check if those numbers are prime. Once everything is successful, it will*
* use RSA encryption to encrypt/decrypt a given text file that the user specify*
*******************************************************************************/
func main(){
    // Prime Number Declarations and number selection
    var num1, num2 int64
    var filename string
    var asciiArray[] uint8

    // Pub/Priv Key Struct Declarations
    var pubKey Keypair
    var privKey Keypair

    // Slice array declarations
    var ciphertext []*big.Int
    var plaintext []string
    var ciph_enc []byte
    


    // Print ASCII Art of the Lecturer and Tutor
    printAsciiArt()

    fmt.Println("Enter a Number (Number Must be between 1000-10000)")
    fmt.Scan(&num1)
    fmt.Println("Enter a Number (Number Must be between 1000-10000) but not the one you entered above")
    fmt.Scan(&num2)

    validateRange(&num1, &num2)

    // Lehmann's Prime Number Check and check if numbers are not the same
    if( (primeCheck(num1)) && (primeCheck(num2)) && (num1 != num2)){
        
        /* Conver num1(p), num2(q) to Big Numbers so we can do calculations */
        p := big.NewInt(num1)
        q := big.NewInt(num2)

        /* Generate KeyPair */
        generate_keypair(p, q, &pubKey, &privKey)

        // The the Public/Private Key Pair
        fmt.Println("Your Public Key: key =", pubKey.key, "n =", pubKey.n)
        fmt.Println("Your Private Key: key =", privKey.key, "n =", privKey.n)

        // Ask for the filename then Read the file
        fmt.Println("Please enter the name of the file you would like to read")
        fmt.Scanf("%s", &filename)
        asciiArray = readFile(filename)

        // Encrypt with your Private Key
        ciphertext = encrypt(asciiArray, privKey)

        // Convert each character of the Cipher text into a String to write to file
        for _, i := range ciphertext{
            // Atoi and Convert to String
            temp, err := strconv.Atoi(i.String())
            // Error handling
            if err != nil{
                fmt.Println(err)
                os.Exit(2)
            }

            // Convert to character and add to my list ready for write out to file
            ciph_enc = append(ciph_enc, byte(temp))
        }

        // Write to File Encrypted Document
        fmt.Println("Please enter the name of the file for writing")
        fmt.Scanf("%s", &filename)
        writeFile(filename, ciph_enc)

        fmt.Println("=====ENCRYPTION FINISHED=====")
        fmt.Println("=====NOW BEGINNING DECRYPTION=====")

        // Decrypt with your public key
        plaintext = decrypt(ciphertext, pubKey)

        // Write your Plaintext to File
        fmt.Println("Please enter the name of the file for writing")
        fmt.Scanf("%s", &filename)
        writePlainText(filename, plaintext)

        fmt.Println("=====ENCRYPTION/DECRYPTION FINISHED=====GOODBYE!!!")
        printNaruto()
    }else{
        fmt.Println("One, or both of the numbers entered is/are not prime or same value...Program Exiting")
    }
}


/*******************************************************************************
* Validates two numbers, ensures that both numbers entered are between         * 
* 1000 and 10 000. One they are there values are exported back to the caller   *
*******************************************************************************/
func validateRange(first *int64, second *int64){
    for (*first < 1000) || (*first > 10000){
        fmt.Println("The First Number is invalid: (Must be between 1000-10000)")
        fmt.Scan(first)
    }

    for (*second < 1000) || (*second > 10000){
        fmt.Println("The Second Number is invalid: (Must be between 1000-10000)")
        fmt.Scan(second)
    }
}


/*******************************************************************************
* Uses Lehmann's algorithm to check if a given value is prime. If it is prime  *
* then the function will use this value as 'p' or 'q' as it's return true. If  *
* the function determines it is not prime by giving a value of either 1 or -1  *
* as the r value then the function will return false.                          *
*******************************************************************************/
func primeCheck(value int64) bool{
    /* Uses r = a^((p-1)/2))modp to determine whether value is a prime number
       , where p is the value being tested and 'a' is just a random number 
       generated that is less than p */
    var a, count int64
    var bigA, bigP *big.Int
    var retVal bool

    // Generate Seed, "UnixNano()" translate the time into milliseconds
    rand.Seed(time.Now().UnixNano())
    
    count = 0
    for (count < 25){
        // Select a random number that is less than Potential Prime
        a = int64(rand.Intn(int(value)))
        
        // Set Big Integer Values
        bigA = big.NewInt(a)
        bigP = big.NewInt(value)
        
        // a^((p-1)/2), calculation
        numerator := big.NewInt(0).Sub(bigP, big.NewInt(1))     // Subtract the Big Numbers, i.e (p-1)
        denominator := big.NewInt(2)                            // Makes 2 a Big Integer
        result := big.NewInt(0).Div(numerator, denominator)     // Uses Integer Division, i.e (p-1)/2
        leftSide := big.NewInt(0).Exp(bigA, result, nil)        // Does the Exponential Calculation
        //fmt.Println("The Left is:", leftSide)

        // left % P calculation
        r := big.NewInt(0).Mod(leftSide, bigP)                  // Mod does the % operation
        //fmt.Println("The Value of r is:", r)

        // Case that the number is 
        if ( (r.Cmp(big.NewInt(1)) == 0) || (r.Cmp(big.NewInt(value - 1)) == 0) ){
            retVal = true
        }else{                                                  // The Number is any other value so it's not prime
            return false
        }
        count++
    }
    
    return retVal
}
