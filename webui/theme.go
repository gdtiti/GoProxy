package webui

const themeBootstrapScript = `<script>
(function () {
  var key = 'goproxy.theme';
  var theme = '';
  try {
    theme = localStorage.getItem(key) || '';
  } catch (e) {}
  if (!theme) {
    try {
      theme = window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches ? 'mist-light' : 'graphite-dark';
    } catch (e) {
      theme = 'graphite-dark';
    }
  }
  document.documentElement.setAttribute('data-theme', theme);
})();
</script>`

const sharedThemeCSS = `
:root{
  color-scheme:dark;
  --bg:#0c1118;
  --bg-elevated:#121923;
  --bg-card:#0f151d;
  --fg:#e5edf5;
  --fg-dim:#97a6b7;
  --fg-text:#d8e1eb;
  --border:#253141;
  --border-heavy:#5eead4;
  --gray-1:#0d1117;
  --gray-2:#17202b;
  --gray-3:#223041;
  --gray-4:#3b4a5c;
  --gray-5:#6d8094;
  --gray-6:#99a9ba;
  --green:#34d399;
  --yellow:#facc15;
  --orange:#fb923c;
  --red:#f87171;
  --mono:JetBrains Mono,Share Tech Mono,monospace;
  --sans:JetBrains Mono,monospace;
  --on-accent:#071411;
  --line-overlay:rgba(94,234,212,0.035);
  --hero-radial:rgba(94,234,212,0.11);
  --surface-glow:rgba(94,234,212,0.14);
  --surface-glow-strong:rgba(94,234,212,0.24);
  --table-hover:rgba(94,234,212,0.06);
  --overlay:rgba(4,8,12,0.74);
  --shadow-soft:0 18px 44px rgba(5,9,14,0.24);
  --shadow-medium:0 24px 60px rgba(5,9,14,0.34);
  --shadow-strong:0 30px 80px rgba(5,9,14,0.4);
  --text-glow:0 0 12px rgba(94,234,212,0.14);
  --logo-glow:0 0 18px rgba(94,234,212,0.18),0 0 36px rgba(94,234,212,0.14);
  --input-ring:0 0 0 4px rgba(94,234,212,0.18);
  --selection-bg:rgba(94,234,212,0.2);
  --muted:#7d8ea3;
  --latency-good:#38bdf8;
  --latency-fair:#a78bfa;
}
:root[data-theme="graphite-dark"]{
  color-scheme:dark;
}
:root[data-theme="aurora-dark"]{
  color-scheme:dark;
  --bg:#08111b;
  --bg-elevated:#0d1825;
  --bg-card:#0b1622;
  --fg:#e8f1ff;
  --fg-dim:#9cb0c9;
  --fg-text:#dce7f6;
  --border:#21364b;
  --border-heavy:#7dd3fc;
  --gray-1:#0a1320;
  --gray-2:#132031;
  --gray-3:#203249;
  --gray-4:#39506b;
  --gray-5:#6b86a3;
  --gray-6:#9fb3c9;
  --green:#38bdf8;
  --yellow:#fde047;
  --orange:#fb923c;
  --red:#f87171;
  --on-accent:#071a29;
  --line-overlay:rgba(125,211,252,0.032);
  --hero-radial:rgba(96,165,250,0.12);
  --surface-glow:rgba(96,165,250,0.14);
  --surface-glow-strong:rgba(96,165,250,0.24);
  --table-hover:rgba(96,165,250,0.07);
  --overlay:rgba(4,8,15,0.76);
  --shadow-soft:0 18px 44px rgba(4,8,15,0.26);
  --shadow-medium:0 24px 60px rgba(4,8,15,0.36);
  --shadow-strong:0 30px 80px rgba(4,8,15,0.42);
  --text-glow:0 0 12px rgba(125,211,252,0.16);
  --logo-glow:0 0 18px rgba(125,211,252,0.2),0 0 36px rgba(96,165,250,0.16);
  --input-ring:0 0 0 4px rgba(125,211,252,0.2);
  --selection-bg:rgba(125,211,252,0.2);
  --latency-good:#22d3ee;
  --latency-fair:#c084fc;
}
:root[data-theme="mist-light"]{
  color-scheme:light;
  --bg:#f4f7fb;
  --bg-elevated:#ffffff;
  --bg-card:#fbfdff;
  --fg:#0f172a;
  --fg-dim:#526277;
  --fg-text:#162033;
  --border:#d7e0ea;
  --border-heavy:#0f766e;
  --gray-1:#ffffff;
  --gray-2:#eef2f7;
  --gray-3:#dce3ec;
  --gray-4:#9aa8b8;
  --gray-5:#64748b;
  --gray-6:#334155;
  --green:#059669;
  --yellow:#ca8a04;
  --orange:#ea580c;
  --red:#dc2626;
  --on-accent:#ffffff;
  --line-overlay:rgba(15,118,110,0.045);
  --hero-radial:rgba(15,118,110,0.11);
  --surface-glow:rgba(15,118,110,0.1);
  --surface-glow-strong:rgba(15,118,110,0.16);
  --table-hover:rgba(15,118,110,0.05);
  --overlay:rgba(15,23,42,0.26);
  --shadow-soft:0 18px 40px rgba(15,23,42,0.08);
  --shadow-medium:0 20px 48px rgba(15,23,42,0.1);
  --shadow-strong:0 28px 64px rgba(15,23,42,0.12);
  --text-glow:none;
  --logo-glow:0 10px 24px rgba(15,118,110,0.08);
  --input-ring:0 0 0 4px rgba(15,118,110,0.12);
  --selection-bg:rgba(15,118,110,0.14);
  --muted:#64748b;
  --latency-good:#0284c7;
  --latency-fair:#7c3aed;
}
:root[data-theme="sand-light"]{
  color-scheme:light;
  --bg:#f5efe6;
  --bg-elevated:#fffaf1;
  --bg-card:#fffdf8;
  --fg:#2d241c;
  --fg-dim:#6d5d4d;
  --fg-text:#372c22;
  --border:#dfcfbb;
  --border-heavy:#8b6f47;
  --gray-1:#fffaf1;
  --gray-2:#efe6d7;
  --gray-3:#dccdb6;
  --gray-4:#b39a7d;
  --gray-5:#876d51;
  --gray-6:#5e4937;
  --green:#4d7c0f;
  --yellow:#ca8a04;
  --orange:#ea580c;
  --red:#b91c1c;
  --on-accent:#fffaf1;
  --line-overlay:rgba(139,111,71,0.05);
  --hero-radial:rgba(139,111,71,0.1);
  --surface-glow:rgba(139,111,71,0.1);
  --surface-glow-strong:rgba(139,111,71,0.16);
  --table-hover:rgba(139,111,71,0.055);
  --overlay:rgba(45,36,28,0.24);
  --shadow-soft:0 18px 40px rgba(69,49,30,0.09);
  --shadow-medium:0 22px 48px rgba(69,49,30,0.11);
  --shadow-strong:0 28px 64px rgba(69,49,30,0.14);
  --text-glow:none;
  --logo-glow:0 12px 28px rgba(139,111,71,0.08);
  --input-ring:0 0 0 4px rgba(139,111,71,0.12);
  --selection-bg:rgba(139,111,71,0.14);
  --muted:#7a6654;
  --latency-good:#0369a1;
  --latency-fair:#7c3aed;
}
html,body{background:var(--bg);color:var(--fg-text)}
body{transition:background-color 0.2s ease,color 0.2s ease}
.theme-switcher{display:inline-flex;align-items:center;gap:10px;padding:8px 12px;border:1px solid var(--border);background:var(--bg-card);box-shadow:var(--shadow-soft);backdrop-filter:blur(10px)}
.theme-label{font-size:10px;letter-spacing:0.08em;text-transform:uppercase;color:var(--fg-dim);font-family:var(--mono);font-weight:600}
.theme-select{padding:8px 12px 8px 12px;min-height:36px;font-size:10px;font-weight:600;cursor:pointer;border:1px solid var(--border);background:var(--bg-elevated);color:var(--fg);font-family:var(--mono);text-transform:uppercase;letter-spacing:0.05em;transition:all 0.2s;outline:none;appearance:none}
.theme-select:hover{border-color:var(--border-heavy);box-shadow:var(--input-ring)}
.theme-select:focus{border-color:var(--border-heavy);box-shadow:var(--input-ring)}
.theme-select option{background:var(--bg-elevated);color:var(--fg)}
::selection{background:var(--selection-bg)}
`

const sharedThemeManagerJS = `
const THEME_STORAGE_KEY = 'goproxy.theme';
const THEME_OPTIONS = [
  { id: 'graphite-dark', labels: { zh: '石墨夜', en: 'Graphite' } },
  { id: 'aurora-dark', labels: { zh: '极夜青', en: 'Aurora' } },
  { id: 'mist-light', labels: { zh: '雾白灰', en: 'Mist' } },
  { id: 'sand-light', labels: { zh: '砂岩米', en: 'Sand' } }
];
let currentTheme = document.documentElement.getAttribute('data-theme') || 'graphite-dark';

function resolveThemeKey(theme) {
  return THEME_OPTIONS.some(function (item) { return item.id === theme; }) ? theme : 'graphite-dark';
}

function detectDefaultTheme() {
  try {
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: light)').matches) return 'mist-light';
  } catch (e) {}
  return 'graphite-dark';
}

function loadThemePreference() {
  let saved = '';
  try {
    saved = localStorage.getItem(THEME_STORAGE_KEY) || '';
  } catch (e) {}
  currentTheme = resolveThemeKey(saved || currentTheme || detectDefaultTheme());
  document.documentElement.setAttribute('data-theme', currentTheme);
  return currentTheme;
}

function applyTheme(theme, persist) {
  currentTheme = resolveThemeKey(theme);
  document.documentElement.setAttribute('data-theme', currentTheme);
  if (persist !== false) {
    try {
      localStorage.setItem(THEME_STORAGE_KEY, currentTheme);
    } catch (e) {}
  }
  const themeSelect = document.getElementById('theme-select');
  if (themeSelect && themeSelect.value !== currentTheme) themeSelect.value = currentTheme;
  window.dispatchEvent(new CustomEvent('goproxy:themechange', { detail: { theme: currentTheme } }));
  return currentTheme;
}

function buildThemeOptions(lang) {
  const locale = lang === 'en' ? 'en' : 'zh';
  return THEME_OPTIONS.map(function (item) {
    return '<option value="' + item.id + '">' + item.labels[locale] + '</option>';
  }).join('');
}

function syncThemeSelect(lang) {
  const themeSelect = document.getElementById('theme-select');
  if (!themeSelect) return;
  themeSelect.innerHTML = buildThemeOptions(lang || 'zh');
  themeSelect.value = currentTheme;
  themeSelect.title = (lang === 'en' ? 'Switch theme' : '切换主题');
}
`
