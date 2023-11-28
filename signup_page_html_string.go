package main

var signupPageHTMLString = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Sticky Labs | Apply</title>
    <link
      href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
      rel="stylesheet"
    />
    <link rel="shortcut icon" href="..ico" type="image/x-icon" />
    <link rel="stylesheet" type="text/css" href="/Sticky-Labs/application.css" />
    <script src="https://unpkg.com/ionicons@5.4.0/dist/ionicons.js"></script>
  </head>
  <body>
    <div class="body_wrap">
      <div class="sidenav">
        <div class="section1">
          <img src="/Sticky-Labs/images/sticky-labs.svg" class="img-fluid" />
        </div>
        <div class="section2">
          <p class="copyright">ALL RIGHTS RESERVED 2023, STICKY LABS.</p>
          <!-- <img src="images/szn.svg" alt="" /> -->

          <svg class="semi">
            <path
              d="M 98,93 L 169,93 C 169,54 137,22 98,22 C 59,22 27,54 27,93 L 98,93 z "
            />
          </svg>

          <!-- C -->

          <svg class="svg-letter" id="svg-1">
            <path
              d="M72.6,146c-9.7,0-19-1.7-27.7-5.2-8.7-3.5-16.4-8.4-23.1-14.8-6.7-6.4-12-14.1-15.9-23C2,94.1,0,84.2,0,73.4v-.4c0-10.4,1.9-20.1,5.7-29,3.8-8.9,9.1-16.7,15.8-23.2,6.7-6.5,14.6-11.6,23.7-15.3C54.3,1.8,64,0,74.4,0s14.8,.9,21.3,2.7c6.5,1.8,12.3,4.3,17.6,7.5,5.3,3.2,9.9,7,14,11.5,4.1,4.5,7.6,9.4,10.5,14.7l-38.6,22.6c-2.7-5.1-6-9.2-9.9-12.3-3.9-3.1-9.1-4.7-15.5-4.7s-7.6,.8-10.8,2.4c-3.2,1.6-5.9,3.8-8.2,6.6-2.3,2.8-4,6.1-5.3,9.8-1.3,3.7-1.9,7.7-1.9,11.8v.4c0,4.5,.6,8.7,1.9,12.5,1.3,3.8,3.1,7.1,5.4,9.8,2.3,2.7,5.1,4.9,8.3,6.4,3.2,1.5,6.7,2.3,10.6,2.3,6.7,0,12-1.6,16-4.9,4-3.3,7.5-7.5,10.4-12.7l38.6,21.8c-2.9,5.2-6.4,10.1-10.4,14.7s-8.7,8.6-14.1,12-11.5,6.1-18.4,8.1c-6.9,2-14.6,3-23.3,3Z"
            />
          </svg>

          <!-- I -->

          <svg class="svg-letter" id="svg-2">
            <path d="M0,0H46.8V140H0V0Z" />
          </svg>

          <!--  -->
          <!-- S -->
          <svg class="svg-letter" id="svg_S">
            <path
              d="M41.6712 171.661C54.198 179.78 66.6087 186.159 78.9035 190.799C91.1983 195.438 102.681 197.758 113.352 197.758C123.559 197.758 131.33 196.366 136.666 193.583C142.001 190.567 144.669 186.159 144.669 180.36C144.669 177.344 143.509 175.024 141.189 173.401C138.87 171.545 135.738 170.037 131.794 168.877C128.083 167.717 123.675 166.789 118.572 166.093C113.468 165.397 108.249 164.469 102.913 163.31C97.8096 162.382 91.7782 161.338 84.8189 160.178C78.0916 158.786 71.1323 156.93 63.941 154.61C56.7497 152.059 49.5584 148.927 42.3672 145.215C35.4079 141.504 29.1445 136.748 23.577 130.949C18.0096 125.149 13.4861 118.306 10.0064 110.419C6.75873 102.3 5.13489 92.7886 5.13489 81.8857C5.13489 65.8793 8.26657 52.6566 14.5299 42.2177C20.7933 31.7787 28.6805 23.5435 38.1916 17.5121C47.9346 11.2488 58.7215 6.9572 70.5524 4.63743C82.3832 2.08568 93.866 0.809799 105.001 0.809799C124.487 0.809799 141.885 3.70951 157.196 9.50894C172.738 15.0764 187.585 23.3116 201.735 34.2145C203.591 35.8383 204.635 37.8101 204.867 40.1299C205.099 42.4496 204.519 44.5374 203.127 46.3933L175.986 85.0174C174.594 86.8732 172.738 88.0331 170.418 88.497C168.099 88.961 166.011 88.497 164.155 87.1052C152.324 79.6819 141.305 74.4624 131.098 71.4467C120.891 68.1991 111.496 66.5752 102.913 66.5752C100.361 66.5752 97.3457 66.6912 93.866 66.9232C90.6183 66.9232 87.4867 67.3871 84.471 68.3151C81.4553 69.011 78.9035 70.2869 76.8157 72.1427C74.9599 73.9985 74.032 76.5502 74.032 79.7979C74.032 83.0456 76.4678 85.9453 81.3393 88.497C86.2108 91.0488 91.7782 93.2526 98.0416 95.1084C104.305 96.9642 110.336 98.472 116.136 99.6319C121.935 100.56 125.647 101.14 127.271 101.372C135.622 102.996 144.785 105.431 154.76 108.679C164.735 111.695 174.014 116.334 182.597 122.598C191.412 128.629 198.72 136.632 204.519 146.607C210.55 156.35 213.566 168.645 213.566 183.492C213.566 197.41 211.13 209.473 206.259 219.68C201.387 229.655 194.544 237.89 185.729 244.385C177.146 250.881 166.707 255.752 154.412 259C142.349 262.016 128.895 263.524 114.048 263.524C101.985 263.524 90.7343 262.48 80.2954 260.392C69.8564 258.536 60.1134 255.868 51.0663 252.389C42.0192 248.909 33.552 244.849 25.6648 240.21C18.0096 235.338 10.9343 230.235 4.43896 224.899C2.58314 223.508 1.53925 221.652 1.30727 219.332C1.07529 216.78 1.65524 214.576 3.0471 212.721L29.8404 174.096C31.2323 172.241 33.0881 171.081 35.4079 170.617C37.7276 170.153 39.8154 170.501 41.6712 171.661Z"
            />
          </svg>

          <!-- T -->
          <svg class="svg-letter" id="svg_T">
            <path
              d="M125.224 42.401H93.517C92.0072 42.401 90.7238 42.9294 89.6669 43.9863C88.61 45.0432 88.0816 46.3265 88.0816 47.8364V161.074C88.0816 162.584 87.5532 163.868 86.4963 164.924C85.5904 165.981 84.307 166.51 82.6462 166.51H48.6748C47.165 166.51 45.8816 165.981 44.8247 164.924C43.7678 163.868 43.2394 162.584 43.2394 161.074V47.8364C43.2394 46.3265 42.7109 45.0432 41.654 43.9863C40.7481 42.9294 39.4648 42.401 37.8039 42.401H5.64437C4.13453 42.401 2.85116 41.8725 1.79428 40.8156C0.73739 39.7587 0.208946 38.4754 0.208946 36.9655V6.16482C0.208946 4.65498 0.73739 3.37161 1.79428 2.31473C2.85116 1.25784 4.13453 0.72939 5.64437 0.72939H125.224C126.884 0.72939 128.168 1.25784 129.074 2.31473C130.131 3.37161 130.659 4.65498 130.659 6.16482V36.9655C130.659 38.4754 130.131 39.7587 129.074 40.8156C128.168 41.8725 126.884 42.401 125.224 42.401Z"
            />
          </svg>

          <!-- I -->
          <!-- <svg class="svg-letter" id="svg_I">
            <path
              d="M61.4678 255H8.92501C6.37326 255 4.28547 254.188 2.66164 252.564C1.26977 250.94 0.573844 248.969 0.573844 246.649V8.98859C0.573844 6.66882 1.26977 4.69702 2.66164 3.07318C4.28547 1.44934 6.37326 0.637429 8.92501 0.637429H61.4678C63.7875 0.637429 65.7593 1.44934 67.3832 3.07318C69.007 4.69702 69.8189 6.66882 69.8189 8.98859V246.649C69.8189 248.969 69.007 250.94 67.3832 252.564C65.7593 254.188 63.7875 255 61.4678 255Z"
            />
          </svg> -->

          <!-- C -->

          <svg class="svg-letter" id="svg-1">
            <path
              d="M72.6,146c-9.7,0-19-1.7-27.7-5.2-8.7-3.5-16.4-8.4-23.1-14.8-6.7-6.4-12-14.1-15.9-23C2,94.1,0,84.2,0,73.4v-.4c0-10.4,1.9-20.1,5.7-29,3.8-8.9,9.1-16.7,15.8-23.2,6.7-6.5,14.6-11.6,23.7-15.3C54.3,1.8,64,0,74.4,0s14.8,.9,21.3,2.7c6.5,1.8,12.3,4.3,17.6,7.5,5.3,3.2,9.9,7,14,11.5,4.1,4.5,7.6,9.4,10.5,14.7l-38.6,22.6c-2.7-5.1-6-9.2-9.9-12.3-3.9-3.1-9.1-4.7-15.5-4.7s-7.6,.8-10.8,2.4c-3.2,1.6-5.9,3.8-8.2,6.6-2.3,2.8-4,6.1-5.3,9.8-1.3,3.7-1.9,7.7-1.9,11.8v.4c0,4.5,.6,8.7,1.9,12.5,1.3,3.8,3.1,7.1,5.4,9.8,2.3,2.7,5.1,4.9,8.3,6.4,3.2,1.5,6.7,2.3,10.6,2.3,6.7,0,12-1.6,16-4.9,4-3.3,7.5-7.5,10.4-12.7l38.6,21.8c-2.9,5.2-6.4,10.1-10.4,14.7s-8.7,8.6-14.1,12-11.5,6.1-18.4,8.1c-6.9,2-14.6,3-23.3,3Z"
            />
          </svg>

          <!-- K -->

          <svg class="svg-letter" id="svg_K">
            <path
              d="M120.842 210.65L63.7507 146.791C62.3774 145.221 61.2003 144.633 60.2193 145.025C59.4346 145.417 59.0422 146.595 59.0422 148.556V208.884C59.0422 210.846 58.3555 212.514 56.9822 213.887C55.6089 215.261 53.9413 215.947 51.9794 215.947H7.83693C5.67885 215.947 3.91315 215.261 2.53983 213.887C1.3627 212.514 0.774135 210.846 0.774135 208.884V7.59486C0.774135 5.63297 1.3627 3.96536 2.53983 2.59205C3.91315 1.21872 5.67885 0.532057 7.83693 0.532057H51.9794C53.9413 0.532057 55.6089 1.21872 56.9822 2.59205C58.3555 3.96536 59.0422 5.63297 59.0422 7.59486V75.28C59.0422 77.2418 59.4346 78.419 60.2193 78.8114C61.0041 79.0075 61.985 78.3209 63.1621 76.7514L116.133 6.41773C117.31 4.84822 118.978 3.57299 121.136 2.59205C123.294 1.41492 125.452 0.826354 127.61 0.826354H183.818C185.976 0.826354 187.35 1.8073 187.938 3.76918C188.527 5.53488 188.232 7.20248 187.055 8.77198L109.659 105.591C108.482 107.161 107.893 109.024 107.893 111.182C108.089 113.144 108.874 114.812 110.247 116.185C116.918 123.444 123.981 130.998 131.436 138.845C138.891 146.693 146.248 154.638 153.507 162.682C160.962 170.53 168.221 178.377 175.284 186.225C182.543 193.876 189.41 201.037 195.884 207.707C197.257 209.277 197.748 211.043 197.355 213.004C196.963 214.966 195.688 215.947 193.53 215.947H132.613C130.455 215.947 128.199 215.457 125.844 214.476C123.686 213.495 122.019 212.22 120.842 210.65Z"
            />
          </svg>

          <!-- Y -->

          <svg class="svg-letter" id="svg_Y">
            <path
              d="M119.684 5.56226L78.7251 71.9149C78.179 72.8979 77.6875 74.0993 77.2506 75.5192C76.8137 76.9391 76.5953 78.1952 76.5953 79.2874V116.805C76.5953 117.898 76.213 118.826 75.4485 119.59C74.6839 120.355 73.7555 120.737 72.6633 120.737H47.9244C46.8322 120.737 45.9038 120.355 45.1392 119.59C44.3747 118.826 43.9924 117.898 43.9924 116.805V79.4512C43.9924 78.359 43.774 77.1029 43.3371 75.6831C43.0094 74.1539 42.5725 72.9525 42.0264 72.0787C35.4731 60.8288 28.7559 49.7973 21.8749 38.9843C15.1031 28.1713 8.33128 17.0306 1.55949 5.56226C1.01338 4.57926 0.904159 3.59625 1.23183 2.61325C1.55949 1.52103 2.26944 0.974918 3.36166 0.974918H32.5241C33.6163 0.974918 34.7085 1.30258 35.8007 1.95792C37.0022 2.61325 37.876 3.43242 38.4221 4.41542L58.5736 39.4758C59.1197 40.3496 59.775 40.7865 60.5396 40.7865C61.3042 40.7865 61.9595 40.3496 62.5056 39.4758L82.4933 4.41542C83.0394 3.43242 83.8586 2.61325 84.9508 1.95792C86.1523 1.30258 87.2991 0.974918 88.3913 0.974918H117.881C118.974 0.974918 119.738 1.35719 120.175 2.12175C120.721 2.88631 120.557 4.03315 119.684 5.56226Z"
            />
          </svg>
        </div>
      </div>

      <div class="content">
        <img src="/Sticky-Labs/images/bg.svg" alt="" class="bg_img" />
        <a class="logout" href="/waitlist/api/logout">
          <ion-icon name="log-out-outline"></ion-icon>
          <span>Logout</span>
        </a>
        <p class="wave2app">WAVE 2 APPLICATION</p>
        <div class="form_wrapper">
          <div class="twitter_welcome">
            <p>Hey <span>@%s</span></p>
            <img src="/Sticky-Labs/images/hand.png" alt="" />
          </div>
          <form action="#" id="form">
            <div class="input-box">
                <input
                    type="text"
                    name=""
                    id="emailInput"
                    placeholder="Enter your email"
                />
                <span id="emailError"></span>
            </div>
            <div class="input-box">
                <input
                    type="text"
                    name=""
                    id="walletInput"
                    placeholder="Enter your wallet address"
                />
                <span id="walletError"></span>
            </div>
            <button class="connectBtn" type="button" onclick="validateAndLog()">Send application</button>
        </form>
        
        
      </div>
    </div>

    <script
      src="https://kit.fontawesome.com/a7f6d58f31.js"
      crossorigin="anonymous"
    ></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/pathseg@1.2.1/pathseg.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/poly-decomp@0.3.0/build/decomp.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/matter-js/0.18.0/matter.min.js"></script>
    <script src="/Sticky-Labs/js/script.js"></script>
  </body>
</html>
`
