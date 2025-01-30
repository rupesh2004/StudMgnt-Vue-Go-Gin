<script>
import { ref } from 'vue';
import axios from "axios";

export default {
    name: "AddStudent",
    setup() {
        const fullName = ref("");
        const Email = ref("");
        const Phone = ref("");
        const dob = ref("");
        const profilePhoto = ref(""); 
        const profilePhotoBase64 = ref(""); 

        const handleProFileChange = (event) => {
            const file = event.target.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onloadend = () => {
                    profilePhotoBase64.value = reader.result.split(',')[1]; 
                    profilePhoto.value = URL.createObjectURL(file); 
                };
                reader.readAsDataURL(file); 
            }
        };

        const createStudent = () => {
            if (fullName.value === "") {
                alert("Please enter your full name");
            } else if (Email.value === "") {
                alert("Please enter your email");
            } else if (Phone.value === "" || Phone.value.length !== 10) {
                alert("Please enter a valid phone number");
            } else if (dob.value === "") {
                alert("Please enter your date of birth");
            } else if (profilePhotoBase64.value === "") {
                alert("Please select a profile photo");
            } else {
                const newStud = {
                    "fullName": fullName.value,
                    "Email": Email.value,
                    "Phone": Phone.value,
                    "dob": dob.value,
                    "profilePhoto": profilePhotoBase64.value 
                };
                axios.post("http://localhost:3000/studData", newStud)
                    .then(response => {
                        console.log("Student added", response.data);
                    })
                    .catch(error => {
                        console.log("Error", error);
                    });
            }
        };

        return {
            fullName,
            Email,
            Phone,
            dob,
            profilePhoto,
            profilePhotoBase64,
            handleProFileChange,
            createStudent,
        };
    }
};
</script scoped>

<template>
    <div class="form-container">
        <h2>Add Student</h2>
        <form @submit.prevent="createStudent">
            <div class="form-group">
                <label>Full Name</label>
                <input type="text" v-model="fullName" />
            </div>
            <div class="form-group">
                <label>Email</label>
                <input type="email" v-model="Email" />
            </div>
            <div class="form-group">
                <label>Phone</label>
                <input type="tel" v-model="Phone" />
            </div>
            <div class="form-group">
                <label>Date of Birth</label>
                <input type="date" v-model="dob" />
            </div>
            <div class="form-group">
                <label>Profile Picture</label>
                <input type="file" @change="handleProFileChange" />
            </div>
            <button type="submit" class="submit-btn">Submit</button>
        </form>
    </div>
</template>

<style scoped>
body {
    font-family: Arial, sans-serif;
    background-color: #f3f4f6;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    margin: 0;
}

.form-container {
    background: #ffffff;
    padding: 25px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
}

h2 {
    text-align: center;
    margin-bottom: 20px;
    color: #333;
}

.form-group {
    margin-bottom: 20px;
}

.form-group label {
    display: block;
    margin-bottom: 8px;
    font-weight: 600;
    color: #444;
}

.form-group input {
    width: 100%;
    padding: 12px 15px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 14px;
    color: #333;
    transition: border-color 0.3s ease;
}

.form-group input:focus {
    outline: none;
    border-color: #4D90FE;
    box-shadow: 0 0 5px rgba(77, 144, 254, 0.5);
}

.submit-btn {
    width: 100%;
    padding: 12px;
    background-color: #4D90FE;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.submit-btn:hover {
    background-color: #357AE8;
}

.submit-btn:active {
    background-color: #2761B6;
}

input[type="file"] {
    padding: 0;
}

input[type="file"]::-webkit-file-upload-button {
    background-color: #4D90FE;
    color: white;
    border: none;
    padding: 8px 15px;
    border-radius: 6px;
    cursor: pointer;
    font-size: 14px;
}

input[type="file"]::-webkit-file-upload-button:hover {
    background-color: #357AE8;
}
</style>
