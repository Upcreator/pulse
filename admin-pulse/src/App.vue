<template>
  <div class="app">
    <form @submit.prevent="login">
      <input type="text" v-model="userInput" class="px-10" placeholder="Enter password" />
      <button type="submit" class="px-10">Логин</button>
    </form>

    <div v-if="isLogin">
      <!-- Filter Controls -->
      <div class="filter-controls">
        <select v-model="filter" @change="filterComments">
          <option value="all">Все пожелания</option>
          <option value="true">Отображаемые пожелания</option>
          <option value="false">Скрытые пожелания</option>
        </select>
      </div>

      <!-- Display Sums -->
      <div v-if="sums.length > 0" class="sums">
        <h2>Суммы</h2>
        <div v-for="sum in sums" :key="sum.id" class="sum-item">
          <p><strong>{{ sum.name }}</strong>: {{ sum.amount }}</p>
          <input 
            type="number" 
            v-model="sum.inputAmount" 
            placeholder="Update amount" 
          />
          <button @click="updateSum(sum.id, sum.name, sum.inputAmount)">Update</button>
        </div>
      </div>

      <!-- Display Comments -->
      <div v-if="filteredComments.length === 0" class="no-comments">
        Нет таких пожеланий.
      </div>

      <CommentCard
        v-for="comment in filteredComments"
        :key="comment.id"
        :comment="comment"
        @updateComment="fetchComments"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import CommentCard from './components/CommentCard.vue';

const userInput = ref(''); // To store the user's input
const isLogin = ref(false); // To manage the login state
const comments = ref([]);
const sums = ref([]);
const filter = ref('all');
const sumIds = ref([
  "816f7317-ee0a-4eb7-801a-deb80d509e65",
  "dc6a0a7d-9ea9-47aa-996e-129ea0de68d0",
  "5f72644b-8807-4c05-aeee-d2129b3c63f1",
  "02046a1e-7710-4bfd-8b4e-fbdd2dbd79ec",
]);

// Login function
const login = () => {
  if (userInput.value === "sdkfn2fndjsf") {
    isLogin.value = true;
    console.log("User logged in successfully");
    fetchComments(); // Fetch comments after login
    fetchSums(); // Fetch sums after login
  } else {
    console.log("Login failed");
  }
};

// Fetch comments from API
const fetchComments = async () => {
  try {
    const response = await axios.get('https://puls-dobra.ru/api/comments');
    comments.value = response.data;
  } catch (error) {
    console.error('Error fetching comments:', error);
  }
};

// Fetch sums from API
const fetchSums = async () => {
  try {
    const promises = sumIds.value.map(id => 
      axios.get(`https://puls-dobra.ru/api/sums/${id}`)
    );
    const responses = await Promise.all(promises);
    sums.value = responses.map(response => ({
      ...response.data,
      inputAmount: response.data.amount // Initialize inputAmount with the current amount
    }));
  } catch (error) {
    console.error('Error fetching sums:', error);
  }
};

// Update sum amount
const updateSum = async (id, name, amount) => {
  try {
    await axios.patch(`https://puls-dobra.ru/api/sums/${id}`, { name, amount });
    console.log(`Sum with ID ${id} updated successfully`);
    fetchSums(); // Refresh the sums after update
  } catch (error) {
    console.error('Error updating sum:', error);
  }
};

// Computed property to filter comments based on the selected filter
const filteredComments = computed(() => {
  if (filter.value === 'all') {
    return comments.value;
  }
  return comments.value.filter(comment => comment.visibility.toString() === filter.value);
});
</script>

<style>
.app {
  padding: 20px;
}

.filter-controls {
  margin-bottom: 20px;
}

.no-comments {
  text-align: center;
  color: #888;
}

.sums {
  margin-top: 20px;
}

.sum-item {
  margin-bottom: 10px;
  display: flex;
  align-items: center;
}

.sum-item input {
  margin-right: 10px;
}
</style>
