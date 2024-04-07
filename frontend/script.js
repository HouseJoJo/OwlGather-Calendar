const btn = document.querySelector(".btn-select");
const items = document.querySelectorAll(".option");

btn.addEventListener("click", ()=>{
  btn.classList.toggle("open");
});

items.forEach(item => {
  item.addEventListener("click", () => {
    item.classList.toggle("checked");
  });
});

const branches = document.querySelector(".btn-select1");
const branchItems = document.querySelectorAll(".option");
branches.addEventListener("click", ()=>{
  branches.classList.toggle("open");
});

const perk = document.querySelector(".btn-select2");
const perkItems = document.querySelectorAll(".option");
perk.addEventListener("click", ()=>{
  perk.classList.toggle("open");
});



