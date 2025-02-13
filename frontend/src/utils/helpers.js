export const formatPrice = (price) => {
    return new Intl.NumberFormat('ru-KZ', {
        style: 'currency',
        currency: "KZT"
    }).format(price);
}