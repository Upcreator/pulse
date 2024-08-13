<!-- src/components/HeroSection.vue -->
<template>
    <section id="hero" class="flex flex-col md:flex-row">
      <div class="hero-image md:w-1/2">
        <img src="/your-image.jpg" alt="Фотография" class="w-full h-auto" />
      </div>
      <div class="hero-text md:w-1/2 p-8 flex items-center">
        <div class="text-content">
          <div class="message-box bg-white shadow-lg rounded-lg p-6">
            <p>Дорогие друзья!</p><br />
            <p>
              Приближается мой день рождения. И я знаю, что многие из вас от
              всего сердца хотят поздравить меня с этим праздником.
            </p><br />
            <p>
              Я всегда с большим трепетом и уважением отношусь к вашему
              вниманию и добрым словам. А ещё – мне очень хочется разделить
              радость этого дня с теми, кто нуждается в поддержке.
            </p><br />
            <p>
              Лучшим подарком для меня станет ваше участие в благотворительной
              акции, которая пройдёт на этом сайте. Если у вас есть желание и
              возможность, пожалуйста, сделайте взнос на благо детей или
              участников специальной военной операции. Общими усилиями мы
              поможем быстрее оплатить чьё-то дорогостоящее лечение, обеспечить
              поддержку защитников Родины и их семей.
            </p><br />
            <p>
              Заранее говорю всем огромное спасибо за понимание, искренность и
              верную дружбу.
            </p><br />
            <p>С уважением,</p>
            <p>Николай Ткаченко.</p>
          </div>
        </div>
      </div>
    </section>
  </template>
  
  <script>
  export default {
    props: {
      isLoading: {
        type: Boolean,
        default: true
      }
    }
  }
  </script>
  
  <style scoped>
  .hero-image img {
    width: 100%;
    height: auto;
  }
  
  .hero-text {
    display: flex;
    align-items: center;
  }
  
  .message-box {
    background-color: #fff;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  
  @media (min-width: 768px) {
    #hero {
      display: flex;
    }
    .hero-image {
      flex: 1;
    }
    .hero-text {
      flex: 1;
      padding: 2rem; 
    }
  }
  </style>
  