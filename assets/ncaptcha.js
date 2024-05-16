(function () {

  let html = ``;

  document.addEventListener("DOMContentLoaded", () => {
    let elements = document.getElementsByClassName("ncaptcha");
    for (let i = 0; i < elements.length; i++) {
      let element = elements.item(i);
      // element.innerHTML = html;
    }


    // click on the checkbox
    elements = document.getElementsByClassName("ncaptcha-checkbox");
    for (let i = 0; i < elements.length; i++) {
      let e = elements.item(i);
      e.addEventListener("click", function () {

        let base = e.closest(".ncaptcha-base");
        let dialog = base.querySelector(".ncaptcha-dialog");
        dialog.showModal();

      })
    }

    elements = document.getElementsByClassName("ncaptcha-dialog");
    for (let i = 0; i < elements.length; i++) {
      let e = elements.item(i);
      e.addEventListener("click", function (e) {
        if (e.target.className === "ncaptcha-dialog") this.close();
      })
    }
    
  });



})()