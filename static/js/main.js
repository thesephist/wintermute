const {
    Component,
} = window.Torus;
const html = window.jdom;

function generate(seed) {
    return fetch(`/generate?seed=${encodeURIComponent(seed)}`)
        .then(r => r.text())
        .catch(e => console.error(e));
}

class App extends Component {
    init() {
        this.seedInput = "the";
        this.generatedText = '';

        this.fetch = this.fetch.bind(this);
    }

    async fetch() {
        this.generatedText = await generate(this.seedInput);
        this.render();
    }

    compose() {
        return html`<main class="app">
            <header class="accent paper">
                <h1>Wintermute</h1>
                <p>
                    Program-generated writing from my
                    <a href="https://thesephist.com/posts/" target="_blank">blog</a>
                </p>
            </header>
            <section class="buttonGroup">
                <div class="left">
                    <input class="paper paper-border-left seedInput" value="${this.seedInput}"
                        placeholder="the"
                        oninput="${evt => {
                            this.seedInput = evt.target.value.trim();
                            this.render();
                        }}"
                        />
                    <button class="movable accent paper generateButton" onclick="${this.fetch}">Generate</button>
                </div>
                <div class="right">
                    <button class="movable paper paper-border-right">
                        <span class="desktop">How does it work</span>?
                    </button>
                </div>
            </section>
            <section class="paper">
                <div class="generatedText">
                    ${this.generatedText ? [
                        ...this.generatedText.split('\n').map(line => {
                            if (line.trim() === '---') {
                                return html`<hr/>`;
                            } else {
                                return html`<p>${line}</p>`;
                            }
                        }),
                     ] : null}
                </div>
                ${this.generatedText ? html`<a class="movable accent paper fullPageButton" target="_blank"
                    href="/full">Generate a full blog post →</a>` : null}
            </section>
            <footer>
                <p>
                    Wintermute powered by
                    <a target="_blank" href="https://github.com/thesephist/torus">Torus</a>
                    ${'&'}
                    <a target="_blank" href="https://thesephist.github.io/paper.css/">paper.css</a>
                </p>
            </footer>
        </main>`;
    }
}

const app = new App();
document.getElementById('root').appendChild(app.node);