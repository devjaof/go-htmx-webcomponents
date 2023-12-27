export class ControlButtons extends HTMLElement {
  public decrease: HTMLButtonElement;

  public increase: HTMLButtonElement;

  constructor() {
    super();
    this.decrease = this.querySelector("#decrease") as HTMLButtonElement;
    this.increase = this.querySelector("#increase") as HTMLButtonElement;
  }

  private handleDisableButton (buttonName: string) {
    if (buttonName == 'Lazer') {
      this.increase.disabled = !this.increase.disabled;
    }

    if (buttonName == 'Produzir') {
        this.decrease.disabled = !this.decrease.disabled;
    }
  }

  private swapOnClick (
    element: HTMLButtonElement,
    buttonName: string,
  ) {
    const that = this;
    element.addEventListener("click",function() {
      if (this.innerText === "Parar") {
        this.innerText = buttonName;
      } else {
        this.innerText = "Parar";
      }

      that.handleDisableButton(buttonName);
    })

  }

  connectedCallback() {
    if (!this.decrease || !this.increase) {
      return;
    }

    this.swapOnClick(this.decrease, "Lazer");
    this.swapOnClick(this.increase, "Produzir");
  }
}

customElements.define("control-buttons", ControlButtons);
