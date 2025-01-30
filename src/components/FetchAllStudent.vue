<script>
import axios from "axios";
import { onMounted, ref } from "vue";

export default {
  name: "FetchAllStudent",
  setup() {
    const students = ref([]);

    const getAllStudentsData = async () => {
      try {
        const response = await axios.get("http://localhost:3000/");
        students.value = response.data.message; 
        console.log(response.data);
      } catch (error) {
        console.error("Error fetching students:", error);
      }
    };

    const deleteStudent = async (studentID) => {
      try {
        await axios.delete(`http://localhost:3000/delete/${studentID}`);
        students.value = students.value.filter((student) => student.studentID !== studentID);
      } catch (error) {
        console.error("Error deleting student:", error);
      }
    };

    const updateStudent = async (studentID) => {
      alert(`Update functionality for Student ID ${studentID} coming soon!`);
    };

    onMounted(() => {
      getAllStudentsData();
    });

    return {
      students,
      deleteStudent,
      updateStudent,
    };
  },
};
</script>

<template>
  <div class="container">
    <h1>Fetching All Students</h1>
    <table class="styled-table">
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
            <img
              :src="student.profilePhoto"
              alt="Profile Photo"
              class="profile-img"
            />
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
  width: 150%;  
  max-width: 1500px;  
  margin: 20px auto;
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
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
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
</style>
