<template>
  <div class="loading-screen">
    <!--<div class="loader mr-4"></div>-->
    <div class="circle-container">
      <img src="/blue-heart.svg" alt="Heart" class="heart-svg"/>
    </div>
    <span class="loading-text">Ваши сердца бьются с нами</span>
  </div>
</template>

<script>
export default {
  name: 'LoadingScreen'
}
</script>

<style scoped>
.loading-screen {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  height: 100vh;
  background-color: #f9f9f9;
}

.loader {
  border: 8px solid #f3f3f3; 
  border-top: 8px solid #3498db; 
  border-radius: 50%;
  width: 50px;
  height: 50px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.circle-container {
  position: relative;
  width: 100px; 
  height: 100px; 
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: radial-gradient(circle, #93c5e6 50%, #fff 100%);
  margin: 10px;
  animation: heartbeat 1s infinite;
}

.heart-svg {
  width: 60px; 
  height: 60px; 
}

@keyframes heartbeat {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}

.loading-text {
  margin-top: 10px;
  font-size: 1.2rem;
  color: #333;
}
</style>
