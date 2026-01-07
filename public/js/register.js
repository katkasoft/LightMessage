const usernameInput = document.getElementById("username");
const emailInput = document.getElementById("email");
const passwordInput = document.getElementById("password");
const confirmInput = document.getElementById("confirm");

function resetBorders() {
    usernameInput.style.borderColor = "";
    emailInput.style.borderColor = "";
    passwordInput.style.borderColor = "";
    confirmInput.style.borderColor = "";
}

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

function register() {
    const username = usernameInput.value;
    const email = emailInput.value;
    const password = passwordInput.value;
    const confirmPass = confirmInput.value;
    
    resetBorders();
    
    let hasError = false;
    
    if (!username) {
        usernameInput.style.borderColor = "red";
        hasError = true;
    }
    
    if (!email || !emailRegex.test(email)) {
        emailInput.style.borderColor = "red";
        hasError = true;
    }
    
    if (!password) {
        passwordInput.style.borderColor = "red";
        hasError = true;
    }
    
    if (!confirmPass || confirmPass !== password) {
        confirmInput.style.borderColor = "red";
        hasError = true;
    }
    
    if (hasError) {
        return;
    }

    registerBtn.disabled = true;
    registerBtn.innerText = "Loading...";

    fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, email, password })
    })
    .then(response => {
        if (!response.ok) {
            registerBtn.disabled = false;
            registerBtn.innerText = "Register";
            
            return response.json().then(errorData => {
                throw new Error(errorData.message || 'Error occurred');
            });
        }
        return response.json();
    })
    .then(data => {
        window.location.href = '/';
    })
    .catch(error => {
        registerBtn.disabled = false;
        registerBtn.innerText = "Register";
        console.error('Error:', error);
        alert(error.message);
    });
}