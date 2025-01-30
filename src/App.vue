<template>
<div class="main-container">
    <h1>Student Management</h1>
    <div class="button-group">
        <button class="action-btn" @click="showAddStudent">Add Student</button>
        <button class="action-btn" @click="getAllStudents">Get All Students</button>
    </div>
    <AddStudent v-if="showForm" />
    <FetchAllStudent v-if="showFetchAllStudForm"/>
</div>
</template>

<script>
import AddStudent from './components/AddStudent.vue';
import FetchAllStudent from './components/FetchAllStudent.vue';

export default {
    name: "App",
    components: {
        AddStudent,
        FetchAllStudent
    },
    data() {
        return {
            showForm: false,
            showFetchAllStudForm: false
        };
    },
    methods: {
        showAddStudent() {
            this.showForm = true;
            this.showFetchAllStudForm = false
        },
        getAllStudents() {
            this.showForm = false
            this.showFetchAllStudForm = true
        }
    }
};
</script>

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

.main-container {
    text-align: center;
    background: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
}

.button-group {
    display: flex;
    justify-content: center;
    gap: 10px;
    margin: 20px 0;
}

.action-btn {
    padding: 10px 15px;
    background: #007bff;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    transition: background 0.3s ease-in-out;
}

.action-btn:hover {
    background: #0056b3;
}
</style>
