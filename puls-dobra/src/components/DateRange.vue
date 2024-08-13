<template>
    <div class="bg-blue-100 p-6 rounded-lg shadow-lg text-center">
        <h2 class="text-xl font-bold text-blue-900">Благотворительная акция</h2>
        <h2 class="text-xl font-bold text-blue-700 pb-8">"Пульс Добра" пройдет:</h2>
        <h2 class="text-xl font-bold text-blue-900">
            {{ startDate }} - {{ endDate }}
        </h2>
    </div>
  </template>
  
  <script setup>
  const startDate = '20 АВГУСТА';
  const endDate = '28 АВГУСТА 2024';
  </script>
  
  <style scoped>
  </style>
  