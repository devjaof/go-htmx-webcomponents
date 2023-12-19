export class ControlButtons extends HTMLElement {
  connectedCallback() {
    this.querySelector("button#decrease")?.addEventListener("click", () => {
      console.log('decrease');
    })
  }
}

customElements.define("control-buttons", ControlButtons);
