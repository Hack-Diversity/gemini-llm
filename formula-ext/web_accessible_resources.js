// This script is injected into the page and is used to add a button to the formula bar
window.addEventListener("load", function (event) {
  if (window.trustedTypes && window.trustedTypes.createPolicy) {
    window.trustedTypes.createPolicy("default", {
      createHTML: (string) => string,
    });
  }

  // Formula bar
  const container = document.querySelector("#t-formula-bar-input-container");
  const item = container.querySelector(".cell-input");

  // Button
  const button = document.createElement("button");
  const innerHTMLText = "Get Formula";
  button.innerHTML =
    window.trustedTypes.defaultPolicy.createHTML(innerHTMLText);

  button.onclick = async function () {
    // add a loading spinner to the button
    button.innerHTML = window.trustedTypes.defaultPolicy.createHTML(
      '<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span> Loading...'
    );
    const text = item.innerText;

    try {
      // make a post request to localhost:8080/formula with the text in the body as "prompt"
      const response = await fetch("http://localhost:8080/formula", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ prompt: text }),
      });

      // Check if response is ok and parse JSON
      if (response.ok) {
        const jsonResponse = await response.json();
        item.innerHTML = window.trustedTypes.defaultPolicy.createHTML(
          jsonResponse["formula"] || "error"
        );

        // press enter to submit the formula
        // const event = new KeyboardEvent("keydown", {
        //   key: "Enter",
        //   code: "Enter",
        //   which: 13,
        //   keyCode: 13,
        //   bubbles: true,
        // });
        // item.dispatchEvent(event);
        button.innerHTML =
          window.trustedTypes.defaultPolicy.createHTML(innerHTMLText);
      } else {
        console.error("Network response was not ok.");
      }
    } catch (error) {
      console.error("There was a problem with the fetch operation:", error);
    }
  };

  //make the button position absolute and blue like the mui button
  button.style.backgroundColor = "#1976d2";
  button.style.color = "white";
  button.style.border = "none";
  button.style.borderRadius = "4px";
  button.style.padding = "5px 10px";
  button.style.cursor = "pointer";
  button.style.position = "absolute";
  button.style.right = "0";
  button.style.top = "0";

  container.appendChild(button);
});
