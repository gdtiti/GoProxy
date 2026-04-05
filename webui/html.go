package webui

var loginHTML = renderLoginHTML("")

var loginHTMLWithError = renderLoginHTML(`<div class="error">[!] ACCESS DENIED - INVALID PASSWORD</div>`)

func renderLoginHTML(errorHTML string) string {
	return `<!DOCTYPE html>
<html lang="zh">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>GoProxy — 身份验证</title>
<link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600;700&family=Share+Tech+Mono&display=swap" rel="stylesheet">
` + themeBootstrapScript + `
<style>
*{box-sizing:border-box;margin:0;padding:0}
` + sharedThemeCSS + `
body{font-family:var(--mono);display:flex;align-items:center;justify-content:center;min-height:100vh;position:relative;padding:32px}
body::before{content:'';position:fixed;inset:0;background:repeating-linear-gradient(0deg,var(--line-overlay) 0px,transparent 1px,transparent 2px,var(--line-overlay) 3px);pointer-events:none;z-index:0}
body::after{content:'';position:fixed;inset:0;background:radial-gradient(ellipse at center,var(--hero-radial) 0%,transparent 72%);pointer-events:none;z-index:0}
.login-shell{position:relative;z-index:1;width:min(100%,960px)}
.topbar{position:fixed;top:20px;right:20px;display:flex;align-items:center;gap:12px;z-index:3}
.card{border:1px solid var(--border-heavy);padding:48px;width:min(100%,460px);background:var(--bg-elevated);box-shadow:var(--shadow-medium);position:relative;z-index:1}
.logo{font-size:58px;margin-bottom:20px;line-height:1;font-weight:700;letter-spacing:0.1em;color:var(--fg);text-shadow:var(--logo-glow)}
h1{font-size:32px;font-weight:700;margin-bottom:10px;letter-spacing:0.15em;text-transform:uppercase;color:var(--fg);text-shadow:var(--text-glow)}
.sub{color:var(--fg-dim);font-size:12px;margin-bottom:38px;letter-spacing:0.08em;text-transform:uppercase}
.field-group{margin-bottom:18px}
label{display:block;font-size:10px;text-transform:uppercase;letter-spacing:0.1em;color:var(--fg-dim);margin-bottom:8px;font-weight:600}
input[type=password]{width:100%;padding:16px;background:var(--bg-card);border:1px solid var(--border);color:var(--fg);font-size:16px;font-family:var(--mono);outline:none;transition:all 0.2s}
input[type=password]:focus{border-color:var(--border-heavy);background:var(--bg-elevated);box-shadow:var(--input-ring)}
button{width:100%;margin-top:24px;padding:16px;background:var(--border-heavy);color:var(--on-accent);border:1px solid var(--border-heavy);font-size:12px;font-weight:700;cursor:pointer;transition:all 0.2s;text-transform:uppercase;letter-spacing:0.1em;font-family:var(--mono);box-shadow:var(--shadow-soft)}
button:hover{transform:translateY(-1px);box-shadow:var(--shadow-medium)}
.error{background:var(--red);color:#fff;padding:14px 16px;font-size:11px;margin-bottom:24px;font-weight:600;border:1px solid var(--red);box-shadow:var(--shadow-soft);text-transform:uppercase;letter-spacing:0.05em}
.tip{color:var(--muted);font-size:11px;margin-top:22px;line-height:1.7;letter-spacing:0.02em}
.tip a{color:var(--border-heavy);text-decoration:none;border-bottom:1px solid transparent;transition:all 0.2s}
.tip a:hover{border-bottom-color:var(--border-heavy)}
.github{display:inline-flex;align-items:center;justify-content:center;width:42px;height:42px;color:var(--fg);background:var(--bg-card);border:1px solid var(--border);opacity:0.92;transition:all 0.2s;box-shadow:var(--shadow-soft)}
.github:hover{transform:translateY(-1px);border-color:var(--border-heavy);box-shadow:var(--shadow-medium)}
.login-theme{min-width:190px}
@media (max-width:640px){
  body{padding:20px}
  .topbar{left:20px;right:20px;justify-content:space-between}
  .login-theme{min-width:0;flex:1}
  .card{padding:32px 24px;margin-top:72px}
  .logo{font-size:48px}
}
</style>
</head>
<body>
<div class="topbar">
  <div class="theme-switcher login-theme">
    <span class="theme-label">主题</span>
    <select id="theme-select" class="theme-select" onchange="applyTheme(this.value)"></select>
  </div>
  <a href="https://github.com/isboyjc/ProxyGo" target="_blank" class="github" title="GitHub">
    <svg width="22" height="22" viewBox="0 0 16 16" fill="currentColor">
      <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/>
    </svg>
  </a>
</div>
<div class="login-shell">
  <div class="card">
    <div class="logo">[GP]</div>
    <h1>GoProxy</h1>
    <p class="sub">// Intelligent Proxy Pool</p>
    ` + errorHTML + `
    <form method="POST" action="/login">
      <div class="field-group">
        <label>> Password</label>
        <input type="password" name="password" placeholder="****************" autofocus>
      </div>
      <button type="submit">[ AUTHENTICATE ]</button>
    </form>
    <p class="tip">访客模式可<a href="/">查看数据</a>，管理员登录后可完全控制。主题选择会保存在当前浏览器中。</p>
  </div>
</div>
<script>
` + sharedThemeManagerJS + `
loadThemePreference();
syncThemeSelect('zh');
</script>
</body>
</html>`
}

// dashboardHTML 已移至 dashboard.go
