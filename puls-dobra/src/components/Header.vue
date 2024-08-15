<!-- Хедер -->
<template>
    <nav>
      <div class="logo-container">
        <div class="logo">Пульс Добра</div>
        <div class="pulse-container">
          <svg viewBox="0 0 1000 100" preserveAspectRatio="none" class="pulse-line-svg">
            <path
              class="pulse-line"
              d="M0 60 L20 60 L40 60 L60 60 L70 60 L80 40 L90 60 L100 55 L110 65 L120 -20 L140 40 L160 40 L170 40 L180 60 L200 60 L220 60 L240 60 L260 60 L280 40 L290 60 L300 55 L310 65 L320 10 L340 40 L370 40 L380 60 L400 60 L420 60 L440 60 L460 60 L480 40 L490 60 L500 55 L510 65 L520 10 L540 40 L560 40 L570 40 L580 60 L600 60 L620 60 L640 60 L660 60 L680 40 L690 60 L700 55 L710 65 L720 10 L740 40 L770 40 L780 60 L800 60 L820 60 L840 60 L860 60 L880 40 L890 60 L900 55 L910 65 L920 10 L940 40 L960 40 L970 40 L980 60 L1000 60"
              stroke-dasharray="1400"
              stroke-dashoffset="1400"
            />
            <defs>
              <filter id="glow">
                <feGaussianBlur in="SourceGraphic" stdDeviation="5" result="blurred" />
                <feMerge>
                  <feMergeNode in="blurred" />
                  <feMergeNode in="SourceGraphic" />
                </feMerge>
              </filter>
            </defs>
          </svg>
        </div>
      </div>
      <div class="nav-links">
        <a href="#charity-section">Акции</a>
        <a href="#carousel">Отчет</a>
      </div>
    </nav>
  </template>
  
  <script>
  export default {
    name: 'Header'
  }
  </script>
  
  <style scoped>
  nav {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 1rem;
    background-color: #3498db;
    color: #fff;
  }
  
  .logo-container {
    display: flex;
    align-items: center;
    margin-bottom: 0.5rem;
  }
  
  .logo {
    font-size: 1.5rem;
    font-weight: bold;
    margin-right: 1rem;
  }
  
  .pulse-container {
    margin: 0;
  }
  
  .pulse-line-svg {
    width: 100%;
    max-width: 200px; 
  }
  
  .pulse-line {
    fill: none;
    stroke: #fff;
    stroke-width: 12;
    animation: pulse 3s linear infinite;
  }
  
  @keyframes pulse {
    0% {
      stroke-dashoffset: 1400;
    }
    99% {
      stroke-dashoffset: 0;
    }
    100% {
      stroke-dashoffset: 1400;
    }
  }
  
  .nav-links {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  
  .nav-links a {
    color: #fff;
    text-decoration: none;
    margin: 0.5rem 0;
  }
  
  .nav-links a:hover {
    text-decoration: underline;
  }
  
  @media (min-width: 768px) {
    nav {
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
    }
  
    .logo-container {
      margin-bottom: 0;
    }
  
    .nav-links {
      flex-direction: row;
    }
  
    .nav-links a {
      margin: 0 1rem;
    }
  }
  </style>
  