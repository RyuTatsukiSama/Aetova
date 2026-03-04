// ─────────────────────────────────────────────
//  LAUNCH BUTTON – plug your game launch logic here
// ─────────────────────────────────────────────
const btnLaunch = document.getElementById('btn-launch');

btnLaunch.addEventListener('click', () => {
  btnLaunch.classList.add('loading');
  btnLaunch.textContent = 'Launching…';

  // TODO: replace this timeout with your actual launch call
  // e.g. window.__TAURI__.tauri.invoke('launch_game') or Electron IPC
  setTimeout(() => {
    btnLaunch.classList.remove('loading');
    btnLaunch.textContent = 'Launch';
  }, 2000);
});