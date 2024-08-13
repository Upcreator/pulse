<template>
  <section id="progress-bar-section">
    <div class="progress-bar-title">
      <h2>Собрано</h2>
    </div>
    <div class="progress-bar-container">
      <div class="progress-bar">
        <div class="progress" :style="{ width: progressWidth }"></div>
      </div>
      <div class="progress-info">
        <span class="amount-collected">{{ formattedCollectedAmount }}</span>
        <span class="amount-target"> / {{ formattedTargetAmount }}</span>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: {
    collectedAmount: {
      type: Number,
      default: 0
    },
    targetAmount: {
      type: Number,
      default: 100000 
    }
  },
  computed: {
    progressWidth() {
      const percentage = (this.collectedAmount / this.targetAmount) * 100;
      return `${Math.min(percentage, 100)}%`; 
    },
    formattedCollectedAmount() {
      return new Intl.NumberFormat().format(this.collectedAmount);
    },
    formattedTargetAmount() {
      return new Intl.NumberFormat().format(this.targetAmount);
    }
  }
}
</script>

<style scoped>
#progress-bar-section {
  padding: 2rem;
  text-align: center;
}

.progress-bar-title {
  margin-bottom: 1rem;
}

.progress-bar-title h2 {
  font-size: 1.8rem;
  font-weight: bold;
  color: #3498db; 
  margin: 0;
}

.progress-bar-container {
  display: inline-block;
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
}

.progress-bar {
  background-color: #e0e0e0;
  border-radius: 8px;
  height: 30px;
  position: relative;
}

.progress {
  background-color: #3498db;
  height: 100%;
  border-radius: 8px;
  transition: width 0.5s ease;
}

.progress-info {
  margin-top: 0.5rem;
  font-size: 1.2rem;
  color: #333;
}

.amount-collected {
  font-weight: bold;
  color: #3498db; 
  font-size: 1.5rem;
}

.amount-target {
  color: #777; 
}
</style>
