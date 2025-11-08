<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'
    import { on } from 'svelte/events';

  async function getQuizzes(){
    try {
      const response = await fetch('http://localhost:3000/api/quizzes')

      if (!response.ok) {
        // non-2xx response
        alert(`Failed to fetch quizzes (status ${response.status})`);
        console.error('Fetch error', response.status, await response.text());
        return;
      }

      const json = await response.json();
      console.log(json);
      alert('Quizzes fetched — check console');
    } catch (err) {
      // network error (CORS, connection refused, DNS, etc.)
      console.error('Network/fetch error', err);
      alert('Failed to fetch quizzes — network error (see console)');
    }
  }
// successful communication with websocket server and client
  function connect(){
    let websocket = new WebSocket("ws://localhost:3000/ws"); //ws mean route
    websocket.onopen = () => {
      console.log("WebSocket connection established");
      websocket.send("Hello Server!");
    };

    websocket.onmessage = (event) => {
      console.log("Message from server ", event.data);
    };
  }
</script>
<button on:click={getQuizzes}>Get Quizzes</button>
<button on:click={connect}>Connect WebSocket</button>
<main>
  <div>
    <a href="https://vite.dev" target="_blank" rel="noreferrer">
      <img src={viteLogo} class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>Vite + Svelte</h1>

  <div class="card">
    <Counter />
  </div>

  <p>
    Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank" rel="noreferrer">SvelteKit</a>, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">
    Click on the Vite and Svelte logos to learn more
  </p>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
