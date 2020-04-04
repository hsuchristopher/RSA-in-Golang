package main

import(
    "fmt"
    "math/big"
    "crypto/rand"
)

/*******************************************************************************
* Once the values of p and q have been established. The function generates the *
* the public/private keypair through running it through the Euclidean's and the*
* extended Euclideans algorithm. But first it needs to calculate the value of  *
* n and totient n. THIS FUNCTION ALSO IMPORTS POINTERS TO THE PUB/PRIV KEYPAIR *
* THEN EXPORTS THE STRUCT VALUES                                               *
*******************************************************************************/
func generate_keypair(p *big.Int, q *big.Int, pubKey *Keypair, privKey *Keypair){
    // Declarations
    var e, gcd, d *big.Int

    //fmt.Println("p and q are:", p, q)
    // n = p*q                                                                      
    n := big.NewInt(0).Mul(p, q)                                                    
    fmt.Println("n is", n)                                                                              

    // phi is the totient of n
    p_minus1 := big.NewInt(0).Sub(p, big.NewInt(1))  
    q_minus1 := big.NewInt(0).Sub(q, big.NewInt(1))
    phi := big.NewInt(0).Mul(p_minus1, q_minus1)
                                                                                    
    // Choose a random integer 'e' so that you can verify using Basic Euclidean 
    /* This uses the crypto/rand library to generate a random number so it is
       more secure than the math/rand library. (It can handle Big Numbers) */

    // Selects a Random Integer that is between 1 and phi (totient)
    e, _ = rand.Int(rand.Reader, phi)

    // Use Basic Euclid's Algorithm to verify that e and phi(n) are coprime
    gcd = basicEuclid(e, phi)
    
    // Keep on selecting a Random number until the GCD is equal to 1
    for gcd.Cmp(big.NewInt(1)) != 0{
        e, _ = rand.Int(rand.Reader, phi)
        gcd = basicEuclid(e, phi)
                                                /* Note to self: Don't implicitly
                                                   ':=' declare variables even if
                                                   have same name outside of loop
                                                   because they will lose scope */
    }
    
    // Use Extended Euclidean's Algorithm to generate the Private Key
    d = extendedEuclid(e, phi)

    // Finally Update your public/private key pairs and export the values
    *pubKey = Keypair{e, n}             // (e, n) is your public Key
    *privKey = Keypair{d, n}            // (d, n) isyour private key
}


/*******************************************************************************
* Finds the Greatest Common Divisor of two integers. Uses Recursion to do this *
*. This rendition of finding the GCD only uses Golang Big Numbers as input     *
*******************************************************************************/
func basicEuclid(a *big.Int, b *big.Int) *big.Int{
    // If 'a' == 0
    if( a.Cmp(big.NewInt(0)) == 0){
        return b
    }
    // Continue the Recursion by b%a and, new value of 'b' is 'a'
    return basicEuclid(big.NewInt(0).Mod(b, a), a)
}

/*******************************************************************************
* Uses the Extended Euclid's algorithm to generate the private key. It returns *
* the modular inverse of the two numbers. For the sake of this algorithm it is *
* random generated number e and phi (totient n)                                *
*******************************************************************************/
func extendedEuclid(a *big.Int, m *big.Int) *big.Int{
    // Declarations
    var m0, x, y, q, t *big.Int

    m0 = m
    y = big.NewInt(0)
    x = big.NewInt(1)

    // if m == 1
    if( m.Cmp(big.NewInt(1)) == 0 ){
        return big.NewInt(0)
    }

    // while a > 1
    for ( a.Cmp(big.NewInt(1)) == 1 ){
        q = big.NewInt(0).Div(a, m)             // Integer Division between 'a' and 'm'
        t = m

        /* m is remainder now, process same as Euclid's algorithm */
        m = big.NewInt(0).Mod(a, m)             // a % m
        a = t
        t = y

        // Update x and y
        y = big.NewInt(0).Mul(q, y)         // y = q * y
        y = big.NewInt(0).Sub(x, y)         // y = x - y
        x = t
    }

    // Make x positive
    // if x < 0
    if ( x.Cmp(big.NewInt(0)) == -1 ){
        x = big.NewInt(0).Add(x, m0)        // x = x + m0
    }
    
    return x
}


/*******************************************************************************
* Takes in the plaintext all as the ASCII decimal value and encrypts each value*
* using the generated Private Key. The function will return the Cipher text    *
* . The algorithm uses each plaintext character to the power of the key % (mod)*
* by 'n' to encrypt each character from plaintext to ciphertext.               *
*******************************************************************************/
func encrypt(plaintext []uint8, key Keypair)[]*big.Int{
    var temp, cipher_character *big.Int

    // Creates an empty nil slice to store cipher text (a.k.a array list thingy)
    var ciphertext []*big.Int

    // For Each Loop
    for  _, i := range plaintext {  // Loops through each character
        // Power it by the Private Key then Mod it by n
        temp = big.NewInt(int64(i))
        
        // Power by the key and Mod by n
        /* The Exp function can take 3 parameters, if the parameters are Exp(x,y,m)
           it is interpretted as x**y % |m| */
        cipher_character = big.NewInt(0).Exp(temp, key.key, key.n)
        //fmt.Println(cipher_character)

        // Append to the end of the list
        ciphertext = append(ciphertext, cipher_character)

    }

    return ciphertext
}


/*******************************************************************************
* Takes in the Cipher text as big in and the public key. The message decrypts  *
* back into the plain text via character to the power of the key % n.          *
* This function returns a String slice of the original plaintext               *
*******************************************************************************/
func decrypt(ciphertext []*big.Int, key Keypair) []string{
    var converted_char string
    var plaintext []string

    // For Each character in the ciphertext
    for _, i := range ciphertext {
        // Power by the key and Mod by n
        /* The Exp function can take 3 parameters, if the parameters are Exp(x,y,m)
           it is interpretted as x**y % |m| */
        plain_char := big.NewInt(0).Exp(i, key.key, key.n)

        // Convert from Big Int to small int64 then back to character
        converted_char = string(plain_char.Int64())


        // Append to the plaintext slice
        plaintext = append(plaintext, converted_char)
    }

    return plaintext
}
