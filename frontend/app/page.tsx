import StatusCard from '@/components/StatusCard';

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center mb-12">
          <h1 className="text-5xl font-bold text-gray-800 mb-4">
            🎓 LabCareer
          </h1>
          <p className="text-xl text-gray-600">
            研究室就活支援プラットフォーム
          </p>
          <p className="text-sm text-gray-500 mt-2">
            Development Environment
          </p>
        </div>

        <div className="flex justify-center">
          <StatusCard />
        </div>

        <div className="text-center mt-12">
          <div className="inline-flex items-center space-x-4 text-sm text-gray-600">
            <span>Phase 0: 開発環境構築</span>
            <span>•</span>
            <span>Docker + Local Hybrid</span>
          </div>
        </div>
      </div>
    </main>
  );
}
