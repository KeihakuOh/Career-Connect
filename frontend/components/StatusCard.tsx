'use client';

import { useEffect, useState } from 'react';
import { healthCheck, dbCheck, apiInfo } from '@/lib/api';

interface Status {
  api: string;
  db: string;
  version: string;
  env: string;
}

export default function StatusCard() {
  const [status, setStatus] = useState<Status>({
    api: '確認中...',
    db: '確認中...',
    version: '-',
    env: '-',
  });

  useEffect(() => {
    const checkStatus = async () => {
      // APIヘルスチェック
      try {
        await healthCheck();
        setStatus(prev => ({ ...prev, api: '✅ 正常' }));
        
        const info = await apiInfo();
        setStatus(prev => ({ 
          ...prev, 
          version: info.version || '-',
          env: info.env || '-'
        }));
      } catch {
        setStatus(prev => ({ ...prev, api: '❌ エラー' }));
      }

      // DBチェック
      try {
        const result = await dbCheck();
        setStatus(prev => ({ ...prev, db: '✅ 接続済み' }));
      } catch {
        setStatus(prev => ({ ...prev, db: '⚠️ 未接続' }));
      }
    };

    checkStatus();
    const interval = setInterval(checkStatus, 10000); // 10秒ごとに更新

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="bg-white rounded-lg shadow-lg p-6 max-w-md w-full">
      <h2 className="text-xl font-bold mb-4">システムステータス</h2>
      
      <div className="space-y-3">
        <div className="flex justify-between items-center py-2 border-b">
          <span className="text-gray-600">API Server</span>
          <span className="font-medium">{status.api}</span>
        </div>
        
        <div className="flex justify-between items-center py-2 border-b">
          <span className="text-gray-600">Database</span>
          <span className="font-medium">{status.db}</span>
        </div>
        
        <div className="flex justify-between items-center py-2 border-b">
          <span className="text-gray-600">Frontend</span>
          <span className="font-medium">✅ 起動中</span>
        </div>
        
        <div className="flex justify-between items-center py-2 border-b">
          <span className="text-gray-600">Version</span>
          <span className="font-mono text-sm">{status.version}</span>
        </div>
        
        <div className="flex justify-between items-center py-2">
          <span className="text-gray-600">Environment</span>
          <span className="font-mono text-sm">{status.env}</span>
        </div>
      </div>

      <div className="mt-6 pt-4 border-t text-xs text-gray-500">
        <div>Frontend: http://localhost:3000</div>
        <div>Backend: http://localhost:8080</div>
        <div>Database: localhost:5432</div>
      </div>
    </div>
  );
}
