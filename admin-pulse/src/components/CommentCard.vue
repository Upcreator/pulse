<template>
  <div :class="cardClass">
    <table class="comment-table">
      <thead>
        <tr>
          <th>Legend</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Автор</td>
          <td>{{ comment.author }}</td>
        </tr>
        <tr>
          <td>Email</td>
          <td>{{ comment.email || 'N/A' }}</td>
        </tr>
        <tr>
          <td>Телефон</td>
          <td>{{ comment.telephone || 'N/A' }}</td>
        </tr>
        <tr>
          <td>Время</td>
          <td>{{ formattedDate }}</td>
        </tr>
        <tr>
          <td>Комментарий</td>
          <td>{{ comment.text }}</td>
        </tr>
      </tbody>
    </table>
    <div class="comment-actions">
      <button @click="acceptComment">Опубликовать</button>
      <button @click="declineComment">Скрыть</button>
      <button class="delete-btn" @click="deleteComment">
        <span class="delete-icon">🗑️</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed, defineEmits } from 'vue';
import axios from 'axios';

const props = defineProps({
  comment: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['updateComment', 'deleteComment']);

const formatDate = (dateString) => {
  const options = { day: '2-digit', month: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit' };
  const date = new Date(dateString);
  return date.toLocaleString('ru-RU', options).replace(',', '');
};

const formattedDate = computed(() => formatDate(props.comment.created_at));

const cardClass = computed(() => {
  return {
    'comment-card': true,
    'visible': props.comment.visibility === true,
    'hidden': props.comment.visibility === false
  };
});

const updateCommentVisibility = async (visibility) => {
  try {
    await axios.patch(`https://puls-dobra.ru/api/comments/${props.comment.id}`, {
      visibility,
      text: props.comment.text,
      author: props.comment.author,
      email: props.comment.email,
      telephone: props.comment.telephone,
      created_at: props.comment.created_at
    });
    emit('updateComment');
    console.log(`${visibility ? 'Accepted' : 'Declined'} comment: ${props.comment.id}`);
  } catch (error) {
    console.error(`${visibility ? 'Error accepting' : 'Error declining'} comment: ${props.comment.id}`, error);
  }
};

const deleteComment = async () => {
  try {
    await axios.delete(`https://puls-dobra.ru/api/comments/${props.comment.id}`);
    emit('deleteComment');
    console.log(`Deleted comment: ${props.comment.id}`);
  } catch (error) {
    console.error(`Error deleting comment: ${props.comment.id}`, error);
  }
};

const acceptComment = () => updateCommentVisibility(true);
const declineComment = () => updateCommentVisibility(false);
</script>

<style scoped>
.comment-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 16px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  max-width: 600px;
  margin: 20px auto;
  transition: background-color 0.3s ease;
}

.comment-card.visible {
  background-color: #e8f5e9; /* Light green */
  border-color: #4CAF50; /* Green */
}

.comment-card.hidden {
  background-color: #fbe9e7; /* Light red */
  border-color: #f44336; /* Red */
}

.comment-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 16px;
}

.comment-table th, .comment-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.comment-table th {
  background-color: #f4f4f4;
  color: #333;
}

.comment-table td {
  color: #444;
}

.comment-actions {
  display: flex;
  align-items: center;
  gap: 10px;
}

.comment-actions button {
  background-color: #4CAF50; /* Green */
  border: none;
  color: white;
  padding: 10px 20px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 0.9em;
  margin: 5px;
  cursor: pointer;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.comment-actions button:hover {
  background-color: #45a049;
}

.comment-actions button:last-child {
  background-color: #f44336;
}

.comment-actions button:last-child:hover {
  background-color: #e53935;
}

.delete-btn {
  background-color: transparent;
  border: none;
  cursor: pointer;
  font-size: 1.2em;
  color: #f44336;
}

.delete-icon {
  display: inline-block;
  vertical-align: middle;
}
</style>
