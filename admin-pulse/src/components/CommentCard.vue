<template>
  <div :class="cardClass">
    <table class="comment-table">
      <thead>
        <tr>
          <th>Автор</th>
          <th>Время</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>{{ comment.author }}</td>
          <td>{{ formattedDate }}</td>
        </tr>
      </tbody>
    </table>
    <p class="comment-text">{{ comment.text }}</p>
    <div class="comment-actions">
      <button @click="acceptComment">Опубликовать</button>
      <button @click="declineComment">Скрыть</button>
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

const emit = defineEmits(['updateComment']);

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
      created_at: props.comment.created_at
    });
    emit('updateComment');
    console.log(`${visibility ? 'Accepted' : 'Declined'} comment: ${props.comment.id}`);
  } catch (error) {
    console.error(`${visibility ? 'Error accepting' : 'Error declining'} comment: ${props.comment.id}`, error);
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

.comment-text {
  font-size: 1em;
  color: #444;
  line-height: 1.5;
  margin-bottom: 16px;
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
  background-color: #f44336; /* Red */
}

.comment-actions button:last-child:hover {
  background-color: #e53935;
}
</style>
