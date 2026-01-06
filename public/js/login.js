const usernameInput = document.getElementById("username");
const passwordInput = document.getElementById("password");

function resetBorders() {
    usernameInput.style.borderColor = "";
    passwordInput.style.borderColor = "";
}

function register() {
    const username = usernameInput.value;
    const password = passwordInput.value;
    
    resetBorders();
    
    let hasError = false
    if (!username) {
        usernameInput.style.borderColor = "red";
        hasError = true;
    }
    if (!password) {
        passwordInput.style.borderColor = "red";
        hasError = true;
    }
    
    if (hasError) {
        return;
    }

    fetch('/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(errorData => {
                throw new Error(errorData.message || 'Error occured');
            });
        }
        return response.json();
    })
    .then(data => {
        setTimeout(() => {
            window.location.href = '/';
        }, 2000);
    })
    .catch(error => {
        console.error('Ошибка:', error);
        alert(error.message || 'Error occured', 'red');
    });
}   