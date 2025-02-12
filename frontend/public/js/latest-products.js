async function fetchProducts() {
    try {
        const response = await fetch("http://127.0.0.1:8080/products?sortField=createdAt&sortOrder=desc&limit=5");
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const products = await response.json();
        console.log(products); // Debugging

        const productContainer = document.getElementById("products-latest");
        productContainer.innerHTML = ""; // Clear previous content

        const row = document.createElement("div");
        row.classList.add("row", "row-cols-5", "g-3"); // 5 columns in a row, with spacing

        products.forEach(product => {
            const productCard = document.createElement("div");
            productCard.classList.add("col");

            const imageUrl = product.LaptopImage?.[0] || "default-image.jpg";
            const cpu = product.Specifications?.CPU || "Unknown";
            const ram = product.Specifications?.RAM || "Unknown";
            const storage = product.Specifications?.Storage || "Unknown";

            productCard.innerHTML = `
                <div class="card h-100 shadow-sm" style="width: 100%;">
                    <img src="${imageUrl}" class="card-img-top" alt="${product.ModelName}" style="height: 200px; object-fit: cover;">
                    <div class="card-body text-center">
                        <h6 class="card-title">${product.ModelName || "Unknown Model"}</h6>
                        <p class="card-text" style="text-align:left">
                            <strong>RAM:</strong> ${ram}<br>
                            <strong>Storage:</strong> ${storage}
                        </p>
                        <a href="#" class="btn btn-dark btn-sm w-100" style="position: relative;">Добавить в корзину</a>
                    </div>
                </div>
            `;

            row.appendChild(productCard);
        });

        productContainer.appendChild(row);
    } catch (error) {
        console.error("Error fetching products:", error);
    }
}

document.addEventListener("DOMContentLoaded", fetchProducts);