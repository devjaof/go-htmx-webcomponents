import {LitElement, css, html} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('control-buttons')
export class ControlButtons extends LitElement {
  static styles = css`p {color: blue}`;

  render() {
    return html`<slot></slot>`;
  }
}

