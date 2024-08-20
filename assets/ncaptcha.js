(function () {

  if (!window.OffscreenCanvas) {
    alert("nCAPTCHA: Your browser does not support OffscreenCanvas. Please update your browser.")
  }
  if (typeof HTMLDialogElement !== 'function') {
    alert("nCAPTCHA: Your browser does not support HTML dialog. Please update your browser.")
  }

  let API = "<PUBLIC_API>";

  document.write(`
    <link href="https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap" rel="stylesheet">
    <link href="${API}/assets/ncaptcha.css" rel="stylesheet">
  `)

  let html = `
<div class="ncaptcha-base">

  <input name="ncaptcha-response" class="ncaptcha-response" style="display: none;">

  <div class="ncaptcha-container">
      <span class="ncaptcha-error"></span>

      <div class="ncaptcha-checkbox"></div>

      <div class="ncaptcha-spinner" style="display: none;"></div>

      <div class="ncaptcha-checkmark" style="display: none;">
          <img src="${API}/assets/checkmark.svg">
      </div>

      <div class="ncaptcha-text">I'm not a human</div>

      <div class="ncaptcha-info">
          <img src="${API}/assets/icon.svg" width="32px" height="32px">
          <div>nCAPTCHA</div>
          <div>Source code</div>
      </div>
  </div>

  <dialog class="ncaptcha-dialog">

      <div class="ncaptcha-question">
          <div class="ncaptcha-title">
              Select all images with
              <strong></strong>
              Click verify once there are none left.
          </div>
          <div class="ncaptcha-images">
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
              <div class="ncaptcha-image"></div>
          </div>
      </div>

      <div class="ncaptcha-actions">
          <button class="ncaptcha-verify-btn" type="button">VERIFY</button>
      </div>

  </dialog>

</div>
  `;

  document.addEventListener("DOMContentLoaded", () => {

    document.querySelectorAll(".ncaptcha").forEach((each) => {
      each.innerHTML = html;

      let checkbox = each.querySelector(".ncaptcha-checkbox");
      let spinner = each.querySelector(".ncaptcha-spinner");
      let checkmark = each.querySelector(".ncaptcha-checkmark");

      let siwtchState = (state) => {
        checkbox.style.display = "none";
        spinner.style.display = "none";
        checkmark.style.display = "none";
        switch (state) {
          case "checkbox":
            checkbox.style.display = "";
            break;
          case "spinner":
            spinner.style.display = "";
            break;
          case "checkmark":
            checkmark.style.display = "";
            break;
        }
      }

      // draw a base64 encoded PNG image to context
      let drawImage = (ctx, img, x, y) => {
        return new Promise((resolve) => {
          let i = new Image();
          i.onload = () => {
            ctx.drawImage(i, x, y, 200, 200, 0, 0, 200, 200);
            resolve()
          }
          i.src = "data:image/png;base64," + img;
        })
      }

      let dialog = each.querySelector(".ncaptcha-dialog");

      checkbox.addEventListener("click", async () => {
        siwtchState("spinner");

        let challengeId = "";
        let answers = [];

        // clear error message
        each.querySelector(".ncaptcha-error").innerHTML = "";

        // clear all event listners
        let clone = dialog.cloneNode(true);
        dialog.parentElement.replaceChild(clone, dialog);
        dialog = clone;

        // enable the button
        dialog.querySelector(".ncaptcha-verify-btn").removeAttribute("disabled");

        // close the dialog when clicking outside
        dialog.addEventListener("click", (e) => {
          if (e.target.className === "ncaptcha-dialog")
            dialog.close();
        })

        // obtain challenge
        let resp = await (await fetch(API + "/challenge")).json();

        challengeId = resp["id"];
        dialog.querySelector(".ncaptcha-title>strong").textContent = resp["select"];

        // split the image and draw to tiles
        let canvas = new OffscreenCanvas(200, 200);
        let context = canvas.getContext("2d");

        let coordinates = [
          [0, 0], [200, 0], [400, 0],
          [0, 200], [200, 200], [400, 200],
          [0, 400], [200, 400], [400, 400]
        ];

        let tiles = dialog.querySelectorAll(".ncaptcha-image");
        for (let i = 0; i < tiles.length; i++) {
          await drawImage(context, resp["challenge"], coordinates[i][0], coordinates[i][1])
          let blob = await canvas.convertToBlob();
          tiles.item(i).style.background = "url(" + URL.createObjectURL(blob) + ")";
          tiles.item(i).style["background-size"] = "cover";

          // reset selected tiles
          tiles.item(i).innerHTML = "";
          tiles.item(i).style.transform = "";
        }

        // click to select a tile
        for (let i = 0; i < tiles.length; i++) {
          tiles.item(i).addEventListener("click", (e) => {
            if (answers.includes(i)) {
              // unselect
              e.target.style.transform = "";
              e.target.innerHTML = "";
              answers = answers.filter((n) => n !== i);
            } else {
              // select
              e.target.style.transform = "scale(0.8)";
              setTimeout(() => {
                e.target.innerHTML = `<img class="ncaptcha-select" src="${API}/assets/checkmark-circle.svg">`;
              }, 100);
              answers.push(i);
            }
          });
        }

        // submit answer
        dialog.querySelector(".ncaptcha-verify-btn").addEventListener("click", async () => {
          dialog.querySelector(".ncaptcha-verify-btn").setAttribute("disabled", true);

          answers.sort((a, b) => a - b);

          let body = new FormData();
          body.set("challenge", challengeId);
          body.set("ans", answers.join(","));

          let resp = await (await fetch(API + "/answer", {
            method: "POST",
            body: body
          })).text();

          if (resp.startsWith("TOKEN_")) {
            each.querySelector(".ncaptcha-response").value = resp.substring(6);
            siwtchState("checkmark");
          } else {
            each.querySelector(".ncaptcha-error").textContent = resp;
          }

          dialog.close();

        });

        dialog.showModal();
        siwtchState("checkbox");
      })
    });

  });

})()