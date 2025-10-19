import './app.css';
import { ConfigProvider } from 'antd';
import { AppRouter } from './app-router.tsx';

export function App() {
  return (
    <div className="app">
      <ConfigProvider>
        <AppRouter />
      </ConfigProvider>
    </div>
  );
}
