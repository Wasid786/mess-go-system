function onScanSuccess(qrCodeMessage) {
    console.log("Scanned QR Code:", qrCodeMessage);

    let student_id;

    try {
        const qrData = JSON.parse(qrCodeMessage);
        if (qrData.student_id) {
            student_id = qrData.student_id.toString();
        } else {
            throw new Error("student_id not found in QR code data");
        }
    } catch (e) {
        student_id = qrCodeMessage.trim();
    }

    if (!/^\d{4,10}$/.test(student_id)) {
        alert("Invalid student ID. Please scan a valid QR code.");
        return;
    }

    fetch('http://localhost:8080/scanner', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ student_id: student_id })
    })
    .then(response => {
        if (!response.ok) {
            return response.json().then(err => { throw new Error(err.error || "Unknown error"); });
        }
        return response.json();
    })
    .then(data => {
        alert(data.message || "Meal Allowed: " + (data.meal_type || ""));
        console.log(data.message);
    })
    .catch(error => {
        alert("Error scanning QR Code: " + error.message);
        console.log(error.message);
    });
}

function onScanFailure(error) {
    console.warn(`QR Code Scan Error: ${error}`);
}

let qrScanner = new Html5QrcodeScanner("qr-reader", { fps: 10, qrbox: 250 });
qrScanner.render(onScanSuccess, onScanFailure);
