package typestruct

type plaintext struct {
	Plaintext string `json:"plaintext"`
}

type ciphertext struct {
	Ciphertext string `json:"ciphertext"`
}

type key struct {
	Key string `json:"key"`
}

type encryptedText struct {
	Plaintext  string `json:"plaintext"`
	Ciphertext string `json:"ciphertext"`
	Key        string `json:"key"`
}
