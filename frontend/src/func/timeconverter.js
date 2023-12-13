export const formatDate = (dateString) => {
  const date = new Date(dateString);

  // Yıl, ay ve günü alarak istenen formatta bir tarih oluşturun
  const day = date.getDate().toString().padStart(2, '0');
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const year = date.getFullYear().toString();

  console.log(`${day}/${month}/${year}`); // Kontrol amaçlı konsola dökümanın yeni formatını yazdıralım

  return `${day}/${month}/${year}`; // İstenen formatta tarih döndürün
};