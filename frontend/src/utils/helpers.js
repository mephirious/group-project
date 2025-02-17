export const formatPrice = (price) => {
    return new Intl.NumberFormat('ru-KZ', {
        style: 'currency',
        currency: "KZT"
    }).format(price);
}

export const formatDate = (date) => {    
    date = new Date(date);
    return`${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}.${date.getFullYear()}`;
}