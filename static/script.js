const textarea = document.getElementById('textarea');
const lineNumbers = document.getElementById('line-numbers');

const endpoint = 'ws://' + window.location.host + '/sync';
const ws = new WebSocket(endpoint);

ws.onmessage = event => {
    const text = event.data;
    textarea.value = text;
    const numberOfLines = text.split('\n').length;
    lineNumbers.innerHTML = Array(numberOfLines).fill('<span></span>').join('');
}

ws.onopen = event => {
    console.log('Connected to server');
}

ws.onerror = event => {
    console.log('Error connecting to server');
    alert(event)
}

textarea.addEventListener('keyup', event=> {
    const text = event.target.value;
    const numberOfLines = text.split('\n').length;
    lineNumbers.innerHTML = Array(numberOfLines).fill('<span></span>').join('');

    ws.send(text);
})