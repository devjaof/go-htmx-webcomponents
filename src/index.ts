export class ControlButtons extends HTMLElement {
  connectedCallback() {
    const decrease = this.querySelector("#decrease") as HTMLButtonElement;
    const increase = this.querySelector("#increase") as HTMLButtonElement;

    if (!decrease || !increase) {
      return;
    }

    decrease.addEventListener("click",function() {
      if (this.innerText !== "Parar") {
        this.innerText = "Parar";
      } else {
        this.innerText = "Lazer";
      }
    })

    increase.addEventListener("click",function() {
      if (this.innerText !== "Parar") {
        this.innerText = "Parar";
      } else {
        this.innerText = "Produzir";
      }
    })
  }
}

customElements.define("control-buttons", ControlButtons);
