<script>
import axios from "axios";
import { onMounted, ref } from "vue";

export default {
  name: "FetchAllStudent",
  setup() {
    const students = ref([]);
    const currentStudent = ref(null);
    const isEditing = ref(false);

    // Fetch all students data
    const getAllStudentsData = async () => {
      try {
        const response = await axios.get("http://localhost:3000/");
        students.value = response.data.message;
      } catch (error) {
        console.error("Error fetching students:", error);
      }
    };

    // Delete student by ID
    const deleteStudent = async (studentID) => {
      try {
        const response = await axios.delete(`http://localhost:3000/deleteStudent/${studentID}`);
        if (response.status === 200) {
          getAllStudentsData();
          alert("Student deleted successfully");
        }
      } catch (error) {
        console.error("Error deleting student:", error);
      }
    };

    // Open update form with pre-filled data
    const updateStudent = (studentID) => {
      const student = students.value.find((s) => s.studentID === studentID);
      currentStudent.value = { ...student };
      isEditing.value = true;
    };

    // Handle file upload and convert to Base64
    const handleFileUpload = (event) => {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          currentStudent.value.profilePhoto = reader.result; // Store Base64 string

        };
      }
    };

    const submitUpdate = async () => {
  try {
    const updatedStudent = {
      fullName: currentStudent.value.fullName,
      Email: currentStudent.value.email,
      Phone: currentStudent.value.phone,
      dob: currentStudent.value.dob,
      profilePhoto: currentStudent.value.profilePhoto,
    };

    console.log("🔹 Sending update request:", JSON.stringify(updatedStudent, null, 2));

    // Check if Base64 string is valid
    if (updatedStudent.profilePhoto.includes("data:image")) {
      updatedStudent.profilePhoto = updatedStudent.profilePhoto.split(",")[1]; // Remove prefix
    }

    const response = await axios.put(
      `http://localhost:3000/updateStudent/${currentStudent.value.studentID}`,
      updatedStudent,
      { headers: { "Content-Type": "application/json" } }
    );

    if (response.status === 200) {
      getAllStudentsData();
      isEditing.value = false;
      alert("Student updated successfully");
    }
  } catch (error) {
    console.error("❌ Error updating student:", error.response?.data || error);
  }
};




    onMounted(() => {
      getAllStudentsData();
    });

    return {
      students,
      deleteStudent,
      updateStudent,
      currentStudent,
      isEditing,
      submitUpdate,
      handleFileUpload,

    };
  },
};
</script>

<template>
  <div class="container">
    <!-- Update Form -->
    <div v-if="isEditing" class="update-form">
      <h2>Update Student</h2>
      <form @submit.prevent="submitUpdate">
        <label for="fullName">Full Name:</label>
        <input type="text" id="fullName" v-model="currentStudent.fullName" required />

        <label for="email">Email:</label>
        <input type="email" id="email" v-model="currentStudent.email" required />

        <label for="phone">Phone:</label>
        <input type="text" id="phone" v-model="currentStudent.phone" required />

        <label for="dob">Date of Birth:</label>
        <input type="date" id="dob" v-model="currentStudent.dob" required />

        <label for="profilePhoto">Profile Photo:</label>
        <input type="file" id="profilePhoto" @change="handleFileUpload" />


        <div class="btn-group">
          <button type="submit" class="update-btn">Update</button>
          <button type="button" @click="isEditing = false" class="cancel-btn">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Students Table -->
    <table class="styled-table" v-else>
      <thead>
        <tr>
          <th>Profile Photo</th>
          <th>Student ID</th>
          <th>Full Name</th>
          <th>Email</th>
          <th>Phone</th>
          <th>Date of Birth</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="student in students" :key="student.studentID">
          <td>
            <img :src="student.profilePhoto" alt="Profile Photo" class="profile-img" />
          </td>
          <td>{{ student.studentID }}</td>
          <td>{{ student.fullName }}</td>
          <td>{{ student.email }}</td>
          <td>{{ student.phone }}</td>
          <td>{{ student.dob }}</td>
          <td class="button-group">
            <button class="update-btn" @click="updateStudent(student.studentID)">Update</button>
            <button class="delete-btn" @click="deleteStudent(student.studentID)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.container {
  max-width: 90%;
  margin: 20px auto;
  text-align: center;
}

h1 {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.styled-table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
  box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
}

.styled-table th, .styled-table td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
  text-align: center;
}

.styled-table th {
  background-color: #007bff;
  color: white;
  font-weight: bold;
}

.styled-table tr:hover {
  background-color: #f1f1f1;
}

.profile-img {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
}

.button-group {
  display: flex;
  gap: 10px;
  justify-content: center;
}

.update-btn, .delete-btn {
  padding: 8px 12px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.update-btn {
  background-color: #28a745;
  color: white;
}

.update-btn:hover {
  background-color: #218838;
}

.delete-btn {
  background-color: #dc3545;
  color: white;
}

.delete-btn:hover {
  background-color: #c82333;
}

/* Update Form */
.update-form {
  width: 50%;
  margin: 30px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
  background-color: #f8f9fa;
}

.update-form label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.update-form input {
  width: 100%;
  padding: 10px;
  margin-bottom: 15px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.btn-group {
  display: flex;
  gap: 10px;
  justify-content: center;
}
</style>
