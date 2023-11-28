function percentX(percent) {
  return Math.round((percent / 100) * window.innerWidth);
}

function percentY(percent) {
  return Math.round((percent / 100) * window.innerHeight);
}

function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1) + min);
}

const Engine = Matter.Engine,
  Bodies = Matter.Bodies,
  Body2 = Matter.Body,
  Svg = Matter.Svg,
  Vertices = Matter.Vertices,
  Constraint = Matter.Constraint,
  Composite = Matter.Composite,
  Mouse = Matter.Mouse,
  MouseConstraint = Matter.MouseConstraint,
  Render = Matter.Render,
  Runner = Matter.Runner;

const engine = Engine.create(),
  world = engine.world;

const render = Render.create({
  element: document.querySelector(".error_section"),
  engine: engine,
  options: {
    wireframes: false,
    showInternalEdges: false,
    width: document.querySelector(".error_section").offsetWidth,
    height: document.querySelector(".error_section").offsetHeight,
    background: "transparent",
  },
});

let bodies = [],
  bgColor = "#0A0618";

// boundaries

var ceiling = Bodies.rectangle(percentX(100) / 2, 10, percentX(100), 20, {
  isStatic: true,
});
var floor = Bodies.rectangle(
  percentX(100) / 2,
  document.querySelector(".error_section").offsetHeight - 10,
  percentX(100),
  20,
  { isStatic: true }
);
var rightWall = Bodies.rectangle(
  document.querySelector(".error_section").offsetWidth - 10,
  document.querySelector(".error_section").offsetHeight / 2,
  20,
  document.querySelector(".error_section").offsetHeight,
  { isStatic: true }
);
var leftWall = Bodies.rectangle(
  10,
  document.querySelector(".error_section").offsetHeight / 2,
  20,
  document.querySelector(".error_section").offsetHeight,
  { isStatic: true }
);
ceiling.render.visible = false;
floor.render.visible = false;
rightWall.render.visible = false;
leftWall.render.visible = false;
bodies.push(ceiling);
bodies.push(floor);
bodies.push(rightWall);
bodies.push(leftWall);

// // circles
for (var i = 0; i < 20; i++) {
  let circleX = getRandomInt(
    10,
    document.querySelector(".error_section").offsetWidth - 5
  );
  let circleY = getRandomInt(
    10,
    document.querySelector(".error_section").offsetHeight - 5
  );
  let circleR = getRandomInt(10, 40);
  let randomColor = Math.floor(Math.random() * 16777215).toString(16);
  let newCircle = Bodies.circle(circleX, circleY, circleR);
  newCircle.render.fillStyle = "#" + randomColor;
  bodies.push(newCircle);
}

// add all bodies (boundaries and circles) to the world
Composite.add(world, bodies);

// semicircles
const semiCircles = [
  ...Array(30)
    .fill()
    .map(() => {
      const path = document.querySelector(".semi > path");
      let randomColor = Math.floor(Math.random() * 16777215).toString(16);
      let randomScale = Math.random() / 2 + 0.1;
      const semi = Bodies.fromVertices(
        Math.random() * document.querySelector(".error_section").offsetWidth,
        Math.random() * document.querySelector(".error_section").offsetHeight,
        Svg.pathToVertices(path),
        {
          render: {
            fillStyle: "#" + randomColor,
            strokeStyle: "#" + randomColor,
            lineWidth: 1,
          },
        },
        true
      );
      const scale = randomScale;
      Matter.Body.scale(semi, scale, scale);
      return semi;
    }),
];

// add all semicircles to the world
Composite.add(world, semiCircles);

// SVGs
let vertexSets = [],
  svgOne,
  svgS,
  svgT,
  svgI,
  svgK,
  svgY;
//   svgTwo,
//   svgThree,
//   svgThreeLegOne,
//   svgThreeLegTwo,
//   svgThreeCounter,
//   svgFour,
//   svgFourCounter;

let cX = percentX(20);
let cY = percentY(20);

let iX = percentX(40);
let iY = percentY(30);

let aX = percentX(60);
let aY = percentY(20);

let yX = percentX(50);
let yY = percentY(60);

let aXLegOne = aX - 43;
let aYLegOne = aY + 49;

let aXLegTwo = aX + 43;
let aYLegTwo = aY + 49;

let oX = percentX(80);
let oY = percentY(30);

let letterSizeHorizontal = 0.8,
  letterSizeVertical = 0.8;

// C
$("#svg-1")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgOne = Bodies.fromVertices(
      cX,
      cY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgOne);
  });

// S
$("#svg_S")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgS = Bodies.fromVertices(
      oX,
      oY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgS);
  });

// T
$("#svg_T")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgT = Bodies.fromVertices(
      iX,
      iY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgT);
  });

// I
$("#svg_I")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgI = Bodies.fromVertices(
      aX,
      aY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgI);
  });

// K
$("#svg_K")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgK = Bodies.fromVertices(
      yX,
      yY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgK);
  });

// Y
$("#svg_Y")
  .find("path")
  .each(function (i, path) {
    let randomColor = Math.floor(Math.random() * 16777215).toString(16);
    svgY = Bodies.fromVertices(
      cX,
      cY,
      Vertices.scale(
        Svg.pathToVertices(path, 10),
        letterSizeHorizontal,
        letterSizeVertical
      ),
      {
        render: {
          fillStyle: "#" + randomColor,
          strokeStyle: "#" + randomColor,
          lineWidth: 2,
        },
      },
      true
    );
    vertexSets.push(svgY);
  });

// I
// $("#svg-2")
//   .find("path")
//   .each(function (i, path) {
//     let randomColor = Math.floor(Math.random() * 16777215).toString(16);
//     svgTwo = Bodies.fromVertices(
//       iX,
//       iY,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: "#" + randomColor,
//           strokeStyle: "#" + randomColor,
//           lineWidth: 2,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgTwo);
//   });

// A
// let randomColorLetterA = Math.floor(Math.random() * 16777215).toString(16);
// $("#svg-3")
//   .find("path")
//   .each(function (i, path) {
//     svgThree = Bodies.fromVertices(
//       aX,
//       aY,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: "#" + randomColorLetterA,
//           strokeStyle: "#" + randomColorLetterA,
//           lineWidth: 2,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgThree);
//   });

// $("#svg-3-leg-1")
//   .find("path")
//   .each(function (i, path) {
//     svgThreeLegOne = Bodies.fromVertices(
//       aXLegOne,
//       aYLegOne,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: "#" + randomColorLetterA,
//           strokeStyle: "#" + randomColorLetterA,
//           lineWidth: 2,
//           isStatic: true,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgThreeLegOne);
//   });

// $("#svg-3-leg-2")
//   .find("path")
//   .each(function (i, path) {
//     svgThreeLegTwo = Bodies.fromVertices(
//       aXLegTwo,
//       aYLegTwo,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: "#" + randomColorLetterA,
//           strokeStyle: "#" + randomColorLetterA,
//           lineWidth: 2,
//           isStatic: true,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgThreeLegTwo);
//   });

// $("#svg-3-counter")
//   .find("path")
//   .each(function (i, path) {
//     svgThreeCounter = Bodies.fromVertices(
//       aX,
//       aY,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: bgColor,
//           strokeStyle: bgColor,
//           lineWidth: 2,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgThreeCounter);
//   });

// create compound body for letter "A"
// var compoundBodyA = Body.create({
//   parts: [svgThree, svgThreeLegOne, svgThreeLegTwo, svgThreeCounter],
// });

// O
// $("#svg-4")
//   .find("path")
//   .each(function (i, path) {
//     let randomColor = Math.floor(Math.random() * 16777215).toString(16);
//     svgFour = Bodies.fromVertices(
//       oX,
//       oY,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: "#" + randomColor,
//           strokeStyle: "#" + randomColor,
//           lineWidth: 1,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgFour);
//   });

// $("#svg-4-counter")
//   .find("path")
//   .each(function (i, path) {
//     svgFourCounter = Bodies.fromVertices(
//       oX,
//       oY,
//       Vertices.scale(
//         Svg.pathToVertices(path, 10),
//         letterSizeHorizontal,
//         letterSizeVertical
//       ),
//       {
//         render: {
//           fillStyle: bgColor,
//           strokeStyle: bgColor,
//           lineWidth: 1,
//         },
//       },
//       true
//     );
//     vertexSets.push(svgFourCounter);
//   });

// create compound body for letter "O"
// var compoundBodyO = Body.create({
//   parts: [svgFour, svgFourCounter],
// });

// add A and O compound bodies to the world
// Composite.add(world, [compoundBodyA, compoundBodyO]);

// add all SVGs to the world
Composite.add(world, vertexSets);

// run the renderer
Render.run(render);

// create runner
const runner = Runner.create();

// run the engine
Runner.run(runner, engine);

// gravity
let intervalID;

function changeGravity() {
  if (!intervalID) {
    intervalID = setInterval(setGravity, 3000);
  }
}

let intervalNumber = 1;
function setGravity() {
  if (intervalNumber === 1) {
    world.gravity.y = 0.5;
    world.gravity.x = 0;
    intervalNumber += 1;
  } else if (intervalNumber === 2) {
    world.gravity.y = -0.5;
    world.gravity.x = 0;
    intervalNumber += 1;
  } else if (intervalNumber === 3) {
    world.gravity.x = 0.5;
    world.gravity.y = 0;
    intervalNumber += 1;
  } else {
    world.gravity.x = -0.5;
    world.gravity.y = 0;
    intervalNumber = 1;
  }
}

changeGravity();

// mouse control
let mouse = Mouse.create(render.canvas),
  mouseConstraint = MouseConstraint.create(engine, {
    mouse: mouse,
    constraint: {
      stiffness: 0.2,
      render: {
        visible: false,
      },
    },
  });

Composite.add(world, mouseConstraint);

// Form validation
function validation() {
  let form = document.getElementById("form");
  let email = document.getElementById("email").value;
  let text = document.getElementById("text");
  let pattern = /^[^ ]+@[^ ]+\.[a-z]{2,3}$/;

  if (email.match(pattern)) {
    form.classList.add("valid");
    form.classList.remove("invalid");
    text.innerHTML = "Your Email Address in valid";
    text.style.color = "#5B6159";
    text.style.fontSize = "14px";
    text.style.fontWeight = "500";
    text.style.fontStyle = "normal";
  } else {
    form.classList.remove("valid");
    form.classList.add("invalid");
    text.innerHTML = "Please Enter Valid Email Address";
    text.style.color = "#ff0000";
    text.style.fontSize = "14px";
    text.style.fontWeight = "500";
    text.style.fontStyle = "normal";
  }

  if (email == "") {
    form.classList.remove("valid");
    form.classList.remove("invalid");
    text.innerHTML = "";
    text.style.color = "#00ff00";
  }
}

// Popup
const popup = document.querySelector(".popup");

const field = popup.querySelector(".field"),
  input = field.querySelector("input"),
  copy = field.querySelector(".popup_button");

copy.onclick = () => {
  input.select(); // Select input value
  if (document.execCommand("copy")) {
    // If the selected text is copied
    field.classList.add("active");
    copy.innerText = "Copied";
    setTimeout(() => {
      window.getSelection().removeAllRanges(); // Remove selection from document
      field.classList.remove("active");
      copy.innerText = "Copy";
    }, 3000);
  }
};
