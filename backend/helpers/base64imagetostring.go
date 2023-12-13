package helpers

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

func SaveBase64Image(base64String string, filename string) error {
	// Base64 kodunun doğruluğunu kontrol etmek için ayrıştırma işlemi
	parts := strings.Split(base64String, ",")
	if len(parts) != 2 {
		return fmt.Errorf("base64 kodu geçersiz")
	}

	// Base64 veri kısmını alıp decode işlemi yapma
	imageBase64 := parts[1]
	decoded, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return err
	}

	// Dosyayı oluştururken ve yazma işlemi yaparken hata kontrolü
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Dosya oluşturma hatası: %s\n", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(decoded)
	if err != nil {
		log.Printf("Dosyaya yazma hatası: %s\n", err)
		return err
	}

	return nil
}
