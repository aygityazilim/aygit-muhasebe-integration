#!/bin/bash
# Test akışını başlatan script

echo ">>> Aygıt Muhasebe Integration - Test Akışı Başlatılıyor <<<"

# 1. Klasör ve Dosya Kontrolü
if [ ! -f "cmd/tester/main.go" ]; then
    echo "Hata: cmd/tester/main.go bulunamadı!"
    exit 1
fi

# 2. Go Çalıştır
go run cmd/tester/main.go

echo ">>> Test Akışı Sona Erdi <<<"
