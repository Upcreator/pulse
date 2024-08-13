<template>
  <section id="charity-section">
    <div class="charity-actions">
      <h2>Благотворительные акции</h2>
    </div>
    <div class="gallery-container">
      <div class="photo-frame">
        <img src="/photo1.jpg" alt="Проект 1" />
        <p><strong>Поможем детям.</strong></p>
        <p>
          У ребятишек, которые борются с тяжёлыми заболеваниями, всегда есть шанс на выздоровление.
          Вместе мы сможем подарить его как можно большему числу детей. Все средства будут направлены
          на закупку лекарств и оплату дорогостоящего лечения подопечных Фонда «Подари жизнь».
        </p>
        <div class="progress-bar-container">
          <h3>Собрано:</h3>
          <span class="text-xs">Данные о сумме собранных пожертвований обнавляется один раз в сутки</span>
          <div class="progress-bar">
            <div class="progress" :style="{ width: progressWidth1 }"></div>
          </div>
          <div class="progress-info">
            <span class="amount-collected">{{ formattedCollectedAmount1 }}</span>
            <!--<span class="amount-target"> / {{ formattedTargetAmount1 }}</span>-->
          </div>
          <button @click="openModal" class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-lg px-6 py-3 text-center">Пожертвовать</button>
        </div>
      </div>
      <div class="photo-frame">
        <img src="/photo2.jpg" alt="Проект 2" />
        <p><strong>Поддержим защитников Отечества.</strong></p>
        <p>
          Герои, которые сегодня находятся на передовой, достойны самого лучшего. И мы можем выполнить
          любые их пожелания, сделать максимум для обеспечения безопасности, помочь комфортно обустроить их быт.
          Все средства будут направлены на поддержку защитников Родины.
        </p>
        <div class="progress-bar-container">
          <h3>Собрано:</h3>
          <span class="text-xs">Данные о сумме собранных пожертвований обнавляется один раз в сутки</span>
          <div class="progress-bar">
            <div class="progress" :style="{ width: progressWidth2 }"></div>
          </div>
          <div class="progress-info">
            <span class="amount-collected">{{ formattedCollectedAmount2 }}</span>
            <!--<span class="amount-target"> / {{ formattedTargetAmount2 }}</span>-->
          </div>
          <button @click="openModal" class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-lg px-6 py-3 text-center">Пожертвовать</button>
        </div>
      </div>
    </div>
    <!-- Modal -->
    <div v-if="isModalOpen" class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
        <div class="bg-white p-6 rounded-lg shadow-lg max-w-md w-full relative z-81">
          <button @click="closeModal" class="absolute top-3 right-3 text-gray-500 hover:text-gray-700">
            &times;
          </button>
  
          <!-- Stepper Navigation -->
          <ol class="flex justify-center items-center w-full p-3 space-x-2 text-sm font-medium text-center text-gray-500 bg-white border border-gray-200 rounded-lg shadow-sm dark:text-gray-400 sm:text-base dark:bg-gray-800 dark:border-gray-700 sm:p-4 sm:space-x-4 rtl:space-x-reverse mb-4">
            <li class="flex items-center text-blue-600 dark:text-blue-500">
              <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-blue-600 rounded-full shrink-0 dark:border-blue-500">
                1
              </span>
              <span>Пожертвование</span>
              <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
              </svg>
            </li>
            <li class="flex items-center" :class="{ 'text-blue-600 dark:text-blue-500': currentStep === 2 }">
              <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">
                2
              </span>
              <span>Пожелания</span>
              <svg v-if="currentStep === 2" class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
              </svg>
            </li>
          </ol>
  
          <!-- Form Steps -->
          <div v-if="currentStep === 1">
            <!-- Clipboard Button and Input -->
            <p id="helper-text-explanation" class="mt-2 text-sm text-gray-500 dark:text-gray-400 pb-2">По номерам этих счетов можно пожертвовать любую сумму, а дальше отправить пожелаение</p>
            <span id="default-message">Помощь детям</span>
            <div class="grid grid-cols-8 gap-2 w-full max-w-[23rem] mb-4">
              <label for="donation-account-1" class="sr-only">Label</label>
              <input ref="donationAccount1" id="donation-account-1" type="text" class="col-span-6 bg-gray-50 border border-gray-300 text-gray-500 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-gray-400 dark:focus:ring-blue-500 dark:focus:border-blue-500" value="1234 1234 1234 1234" disabled readonly>
              <button @click="copyToClipboard('donationAccount1')" class="col-span-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 items-center inline-flex justify-center">
                <span id="default-message">Copy</span>
                <span id="success-message" class="hidden inline-flex items-center">
                  <svg class="w-3 h-3 text-white me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"/>
                  </svg>
                  Copied!
                </span>
              </button>
            </div>
            <span id="default-message">Помощь Защитникам Отечества</span>
            <div class="grid grid-cols-8 gap-2 w-full max-w-[23rem] mb-4">
              <label for="donation-account-2" class="sr-only">Label</label>
              <input ref="donationAccount2" id="donation-account-2" type="text" class="col-span-6 bg-gray-50 border border-gray-300 text-gray-500 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-gray-400 dark:focus:ring-blue-500 dark:focus:border-blue-500" value="1234 1234 1234 1234" disabled readonly>
              <button @click="copyToClipboard('donationAccount2')" class="col-span-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 items-center inline-flex justify-center">
                <span id="default-message">Copy</span>
                <span id="success-message" class="hidden inline-flex items-center">
                  <svg class="w-3 h-3 text-white me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"/>
                  </svg>
                  Copied!
                </span>
              </button>
            </div>
            <button @click="nextStep" class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Далее</button>
          </div>
  
          <div v-if="currentStep === 2">
            <form @submit.prevent="handleSubmit" class="max-w-sm mx-auto">
              <div class="mb-5">
                <label for="author" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Ваше Имя</label>
                <input type="author" id="author" v-model="email" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" placeholder="Введите Ваше Имя" required />
              </div>
              <div class="mb-5">
                <label for="telephone" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Введите ваш номер телефона <span class="text-red-500">*</span></label>
                <p id="helper-text-explanation" class="mt-2 text-sm text-gray-500 dark:text-gray-400 pb-2"><span class="text-red-500">*</span> Мне важно сказать тебе спасибо</p>
                <input type="telephone" id="telephone" v-model="telephone" class="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 dark:shadow-sm-light" placeholder="Оставьте Ваш Номер" />
              </div>
              <div class="mb-5">
                <label for="message" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Напишите Ваше пожелание</label>
                <textarea id="message" v-model="message" rows="4" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Ваше Пожелание"></textarea>
              </div>
              <div class="mb-5 flex items-center">
                <label class="inline-flex items-center cursor-pointer">
                  <input type="checkbox" id="terms" v-model="terms" class="sr-only peer" required />
                  <div class="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:w-5 after:h-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600"></div>
                  <span class="ms-3 text-sm font-medium text-gray-900 dark:text-gray-300">Отмечая, Вы даёте <a href="#" class="text-blue-600 hover:underline dark:text-blue-500">согласие с обработкой персональных данных</a></span>
                </label>
              </div>
              <div class="flex justify-between">
                <button @click="prevStep" type="button" class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Назад</button>
                <button type="submit" class="text-white bg-blue-500 hover:bg-blue-600 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Отправить</button>
              </div>
            </form>
          </div>
        </div>
      </div>
  </section>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      collectedAmount1: 0, // Сообранная сумма для Детей
      targetAmount1: 100000,   // Целевая сумма для Детей
      collectedAmount2: 0, // Собранная сумма для СВО
      targetAmount2: 70000,    // Целевая сумма для СВО
      isModalOpen: false,
      currentStep: 1, 
      author: '',
      telephone: '',
      message: '',
      terms: false
    };
  },
  methods: {
    openModal() {
      this.isModalOpen = true;
    },
    closeModal() {
      this.isModalOpen = false;
      this.currentStep = 1; 
    },
    nextStep() {
      if (this.currentStep < 2) {
        this.currentStep++;
      }
    },
    prevStep() {
      if (this.currentStep > 1) {
        this.currentStep--;
      }
    },
    async handleSubmit() {
      try {
        const response = await axios.post('http://87.228.19.168:8080/comments', {
          author: this.author,
          text: this.message,
          visibility: false
        });
        console.log(response);
        this.closeModal(); 
      } catch (error) {
        console.error('Error submitting form:', error);
        alert('There was an error submitting the form. Please try again.');
      }
    },
    copyCardNumber(cardNumber) {
      navigator.clipboard.writeText(cardNumber).then(() => {
        alert('Номер счёта скопирован!');
      });
    }
  },
  computed: {
    progressWidth1() {
      const percentage = (this.collectedAmount1 / this.targetAmount1) * 100;
      return `${Math.min(percentage, 100)}%`;
    },
    progressWidth2() {
      const percentage = (this.collectedAmount2 / this.targetAmount2) * 100;
      return `${Math.min(percentage, 100)}%`;
    },
    formattedCollectedAmount1() {
      return new Intl.NumberFormat().format(this.collectedAmount1);
    },
    formattedTargetAmount1() {
      return new Intl.NumberFormat().format(this.targetAmount1);
    },
    formattedCollectedAmount2() {
      return new Intl.NumberFormat().format(this.collectedAmount2);
    },
    formattedTargetAmount2() {
      return new Intl.NumberFormat().format(this.targetAmount2);
    }
  }
}
</script>
<style scoped>
#charity-section {
  padding: 2rem;
  background-color: #f9f9f9;
  text-align: center;
}

.charity-actions h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: #333;
}

.gallery-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 1rem;
}

.photo-frame {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.2);
  overflow: hidden;
  padding: 1rem;
  width: 100%;
  max-width: 400px;
  text-align: left; 
}

.photo-frame img {
  width: 100%;
  height: auto;
  display: block;
}

.progress-bar-container {
  text-align: center; 
  margin-top: 1rem;
}

.progress-bar {
  background-color: #e0e0e0;
  border-radius: 8px;
  height: 30px;
  position: relative;
  margin: 0 auto; 
  width: 100%; 
  max-width: 300px;
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

.progress-bar-container h3 {
  font-size: 1.2rem;
  font-weight: bold;
  color: #3498db; 
}

.amount-collected {
  font-weight: bold;
  color: #3498db;
  font-size: 1.5rem;
}

.amount-target {
  color: #777;
}

.card-number {
  margin-top: 1rem;
  color: #3498db;
  cursor: pointer;
  text-decoration: underline;
}

@media (min-width: 768px) {
  .gallery-container {
    flex-direction: row;
    justify-content: center;
  }

  .photo-frame {
    width: calc(50% - 0.5rem);
  }
}
</style>
