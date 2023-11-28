// Disable Scroll for few seconds
$(function () {
  setTimeout(function () {
    $("html, body").css({
      "overflow-y": "auto",
    });
  }, 4000);
  return false;
});

const svg = document.querySelector("#svg");
const mouse = svg.createSVGPoint();

const leftEye = createEye("#left_eye");
const rightEye = createEye("#right_eye");

const leftEye2 = createEye("#left_eye2");
const rightEye2 = createEye("#right_eye2");

let requestId = null;

window.addEventListener("mousemove", onMouseMove);

function onFrame() {
  let point = mouse.matrixTransform(svg.getScreenCTM().inverse());
  let point2 = mouse.matrixTransform(svg.getScreenCTM());

  leftEye.rotateTo(point);
  rightEye.rotateTo(point);
  leftEye2.rotateTo(point);
  rightEye2.rotateTo(point);

  requestId = null;
}

function onMouseMove(event) {
  mouse.x = event.clientX;
  mouse.y = event.clientY;

  if (!requestId) {
    requestId = requestAnimationFrame(onFrame);
  }
}

function createEye(selector) {
  const element = document.querySelector(selector);

  TweenMax.set(element, {
    transformOrigin: "center",
  });

  let bbox = element.getBBox();
  let centerX = bbox.x + bbox.width / 2;
  let centerY = bbox.y + bbox.height / 2;

  function rotateTo(point) {
    let dx = point.x - centerX;
    let dy = point.y - centerY;

    let angle = Math.atan2(dy, dx);

    TweenMax.to(element, 0.3, {
      rotation: angle + "_rad_short",
    });
  }

  return {
    element,
    rotateTo,
  };
}

const cursor = document.querySelector("#cursor");
const cursorCircle = cursor.querySelector(".cursor__circle");

const mouse2 = { x: -100, y: -100 }; // mouse pointer's coordinates
const pos = { x: 0, y: 0 }; // cursor's coordinates
const speed = 0.1; // between 0 and 1

const updateCoordinates = (e) => {
  mouse2.x = e.clientX;
  mouse2.y = e.clientY;
};

window.addEventListener("mousemove", updateCoordinates);

function getAngle(diffX, diffY) {
  return (Math.atan2(diffY, diffX) * 180) / Math.PI;
}

function getSqueeze(diffX, diffY) {
  const distance = Math.sqrt(Math.pow(diffX, 2) + Math.pow(diffY, 2));
  const maxSqueeze = 0.15;
  const accelerator = 1500;
  return Math.min(distance / accelerator, maxSqueeze);
}

const updateCursor = () => {
  const diffX = Math.round(mouse2.x - pos.x);
  const diffY = Math.round(mouse2.y - pos.y);

  pos.x += diffX * speed;
  pos.y += diffY * speed;

  const angle = getAngle(diffX, diffY);
  const squeeze = getSqueeze(diffX, diffY);

  const scale = "scale(" + (1 + squeeze) + ", " + (1 - squeeze) + ")";
  const rotate = "rotate(" + angle + "deg)";
  const translate = "translate3d(" + pos.x + "px ," + pos.y + "px, 0)";

  cursor.style.transform = translate;
  cursorCircle.style.transform = rotate + scale;
};

function loop() {
  updateCursor();
  requestAnimationFrame(loop);
}

requestAnimationFrame(loop);

const cursorModifiers = document.querySelectorAll("[cursor-class]");

cursorModifiers.forEach((curosrModifier) => {
  curosrModifier.addEventListener("mouseenter", function () {
    const className = this.getAttribute("cursor-class");
    cursor.classList.add(className);
  });

  curosrModifier.addEventListener("mouseleave", function () {
    const className = this.getAttribute("cursor-class");
    cursor.classList.remove(className);
  });
});
