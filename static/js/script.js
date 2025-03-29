function onScanSuccess(qrCodeMessage) {
    console.log("Scanned QR Code:", qrCodeMessage);

    let course_id;

    try {
        const qrData = JSON.parse(qrCodeMessage);
        if (qrData.course_id) {
            course_id = qrData.course_id.toString();
        } else {
            throw new Error("course_id not found in QR code data");
        }
    } catch (e) {
        course_id = qrCodeMessage.trim();
    }

    fetch('http://localhost:8080/scan', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ course_id: course_id })
    })
    .then(async response => {
        if (!response.ok) {
            const err = await response.json();
            throw new Error(err.error || "Unknown error");
        }
        console.log(response.json());
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
