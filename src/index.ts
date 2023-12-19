export class ControlButtons extends HTMLElement {
  private swapOnClick(element: HTMLButtonElement, buttonName: string) {
    element.addEventListener("click",function() {
      if (this.innerText !== "Parar") {
        this.innerText = "Parar";
      } else {
        this.innerText = buttonName;
      }
    })
  }
  
  connectedCallback() {
    const decrease = this.querySelector("#decrease") as HTMLButtonElement;
    const increase = this.querySelector("#increase") as HTMLButtonElement;

    if (!decrease || !increase) {
      return;
    }

    this.swapOnClick(decrease, "Lazer");
    this.swapOnClick(increase, "Produzir");
  }
}

customElements.define("control-buttons", ControlButtons);
