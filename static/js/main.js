const {
    Component,
} = window.Torus;
const html = window.jdom;

const FAUX_LOADING_TIMEOUT = 700;

function generate(seed) {
    return fetch(`/generate?seed=${encodeURIComponent(seed)}`)
        .then(r => r.text())
        .catch(e => console.error(e));
}

function Popup(children) {
    return html`<div class="popup-wrapper">
        <div class="paper paper-border-top popup">${children}</div>
    </div>`;
}

class App extends Component {
    init() {
        this.seedInput = "when";
        this.generatedText = '';

        this._isLoading = false;
        this._isShowingHow = false;

        this.fetch = this.fetch.bind(this);
        this.handleInput = this.handleInput.bind(this);

        this.fetch();
    }

    fetch() {
        this._isLoading = true;
        this.render();

        setTimeout(async () => {
            this.generatedText = (await generate(this.seedInput)).trim() + '...';
            this._isLoading = false;
            this.render();
        }, FAUX_LOADING_TIMEOUT)
    }

    handleInput(evt) {
        this.seedInput = evt.target.value.trim();
        this.render();
    }

    compose() {
        return html`<main class="app">
            <header class="accent paper">
                <h1>Wintermute</h1>
                <p>
                    Program-generated writing from my
                    <a href="https://thesephist.com/posts/" target="_blank">blog</a>.
                </p>
            </header>
            <section class="buttonGroup">
                <div class="left">
                    <input class="paper paper-border-left seedInput" value="${this.seedInput}"
                        placeholder="the"
                        oninput="${this.handleInput}"
                        onkeydown="${evt => {
                            if (evt.key == 'Enter') {
                                this.fetch();
                            }
                        }}"
                        />
                    <button class="movable accent paper generateButton" onclick="${this.fetch}">Generate</button>
                </div>
                <div class="right">
                    <button class="movable paper paper-border-right"
                        onclick="${() => {
                            this._isShowingHow = true;
                            this.render();

                            const closeButton = this.node.querySelector('.closeButton');
                            closeButton && closeButton.focus();
                        }}">
                        <span class="desktop">How does it work</span>?
                    </button>
                </div>
            </section>
            <section class="paper">
                <div class="generatedText">
                    ${!this._isLoading && this.generatedText ? this.generatedText.split('\n').map(line => {
                        if (line.trim() === '---') {
                            return html`<hr/>`;
                        } else {
                            return html`<p>${line}</p>`;
                        }
                    }) : html`<div class="loading">
                        <div class="accent paper square first"/>
                        <div class="accent paper square second"/>
                        <div class="accent paper square third"/>
                    </div>`}
                </div>
                ${!this._isLoading && this.generatedText ? html`<a class="movable accent paper fullPageButton" target="_blank"
                    href="/full">Generate a full blog post →</a>` : null}
            </section>
            <footer>
                <p>
                    Wintermute is powered by
                    <a target="_blank" href="https://github.com/thesephist/torus">Torus</a>
                    ${'&'}
                    <a target="_blank" href="https://thesephist.github.io/paper.css/">paper.css</a>
                </p>
            </footer>
            ${this._isShowingHow ? Popup(html`<div class="howItWorks">
                <p><strong>How does Wintermute work?</strong></p>
                <p>
                    Wintermute creates fake blog post writing with a
                    <a href="https://en.wikipedia.org/wiki/Markov_chain" target="_blank">Markov Chain</a>
                    based on my archive of blog posts.
                </p>
                <div class="buttonGroup">
                    <div class="left"></div>
                    <div class="right">
                        <a href="https://github.com/thesephist/wintermute" target="_blank"
                            class="movable paper">See on GitHub</a>
                        <button class="movable paper closeButton"
                            onclick="${() => {
                                this._isShowingHow = false;
                                this.render();
                            }}">Close</button>
                    </div>
                </div>
            </div>`) : null}
        </main>`;
    }
}

const app = new App();
document.getElementById('root').appendChild(app.node);