document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");

    form.addEventListener("submit", async function (event) {
        event.preventDefault(); 

        const formData = {
            name: document.getElementById("name").value,
            hostel: document.getElementById("hostel").value,
            course_id: document.getElementById("course_id").value,
            enroll_no: document.getElementById("enroll_no").value,
            registered_session: document.getElementById("registered_session").value,
            mess_slip: document.getElementById("mess_slip").value,
        };

        try {
            const response = await fetch("/students", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(formData),
            });

            const result = await response.json();

            if (response.ok) {
                alert("Student Registered Successfully!");
                form.reset();
            } else {
                alert("Error: " + result.message);
            }
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("An error occurred while submitting the form.");
        }
    });
});
