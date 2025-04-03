import { useState } from "react";
import { Command, CommandInput } from "@/components/ui/command";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Separator } from "@/components/ui/separator";
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  Terminal,
  Code,
  Play,
  Download,
  BookOpen,
  Settings,
  History,
  ArrowRight,
} from "lucide-react";

export default function App() {
  const [commandInput, setCommandInput] = useState("");
  const [commandHistory, setCommandHistory] = useState<string[]>([]);
  const [output, setOutput] = useState("");

  const commonCommands = [
    { name: "ls", description: "Listar arquivos" },
    { name: "cd", description: "Mudar diretório" },
    { name: "mkdir", description: "Criar diretório" },
    { name: "rm", description: "Remover arquivo" },
    { name: "git status", description: "Ver status do git" },
    { name: "npm install", description: "Instalar dependências" },
  ];

  const executeCommand = () => {
    if (!commandInput.trim()) return;

    const newHistory = [...commandHistory, commandInput];
    setCommandHistory(newHistory);

    // Simule a saída do comando
    setOutput(
      `Executando: ${commandInput}\n\n> Comando executado com sucesso!`
    );
    setCommandInput("");
  };

  const selectCommand = (cmd: string) => {
    setCommandInput(cmd);
  };

  return (
    <div className="min-h-screen bg-gray-950 text-gray-200 flex flex-col">
      {/* Barra superior */}
      <header className="bg-gray-900 p-4 border-b border-gray-800 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Terminal className="h-6 w-6 text-indigo-400" />
          <h1 className="text-xl font-bold">Terminal GUI</h1>
        </div>
        <div className="flex gap-2">
          <Button variant="ghost" size="icon">
            <Settings className="h-5 w-5" />
          </Button>
          <Button variant="ghost" size="icon">
            <BookOpen className="h-5 w-5" />
          </Button>
        </div>
      </header>

      {/* Conteúdo principal */}
      <div className="flex flex-1 overflow-hidden">
        {/* Sidebar */}
        <div className="w-64 bg-gray-900 border-r border-gray-800 flex flex-col">
          <div className="p-4">
            <h2 className="text-lg font-medium mb-2">Comandos Recentes</h2>
            <ScrollArea className="h-64">
              {commandHistory.length > 0 ? (
                <ul className="space-y-1">
                  {commandHistory.map((cmd, index) => (
                    <li
                      key={index}
                      className="px-2 py-1 hover:bg-gray-800 rounded cursor-pointer flex items-center"
                      onClick={() => selectCommand(cmd)}
                    >
                      <History className="h-4 w-4 mr-2 text-gray-400" />
                      <span className="truncate">{cmd}</span>
                    </li>
                  ))}
                </ul>
              ) : (
                <p className="text-gray-500 text-sm p-2">
                  Nenhum comando executado
                </p>
              )}
            </ScrollArea>
          </div>

          <Separator className="bg-gray-800" />

          <div className="p-4 flex-1">
            <h2 className="text-lg font-medium mb-2">Documentação</h2>
            <div className="text-sm text-gray-400">
              <p className="mb-2">
                Use os comandos com os parâmetros apropriados.
              </p>
              <p>Exemplos:</p>
              <ul className="space-y-1 mt-1">
                <li>
                  <code className="text-indigo-400">ls -la</code>
                </li>
                <li>
                  <code className="text-indigo-400">
                    git commit -m "mensagem"
                  </code>
                </li>
              </ul>
            </div>
          </div>

          <div className="p-4 border-t border-gray-800">
            <Button
              variant="outline"
              className="w-full flex items-center gap-2"
            >
              <Download className="h-4 w-4" />
              Exportar Logs
            </Button>
          </div>
        </div>

        {/* Área principal */}
        <div className="flex-1 flex flex-col">
          <Tabs defaultValue="terminal" className="flex-1 flex flex-col">
            <div className="bg-gray-900 p-2 border-b border-gray-800">
              <TabsList className="bg-gray-800">
                <TabsTrigger
                  value="terminal"
                  className="flex items-center gap-2"
                >
                  <Terminal className="h-4 w-4" />
                  Terminal
                </TabsTrigger>
                <TabsTrigger
                  value="commands"
                  className="flex items-center gap-2"
                >
                  <Code className="h-4 w-4" />
                  Comandos
                </TabsTrigger>
              </TabsList>
            </div>

            <TabsContent
              value="terminal"
              className="flex-1 flex flex-col p-0 m-0"
            >
              <ScrollArea className="flex-1 p-4 bg-gray-950">
                <div className="font-mono whitespace-pre-wrap">
                  {output || (
                    <span className="text-gray-500">
                      A saída do comando aparecerá aqui...
                    </span>
                  )}
                </div>
              </ScrollArea>

              <div className="p-4 bg-gray-900 border-t border-gray-800 flex items-center gap-2">
                <div className="flex-1">
                  <Command className="rounded-lg border border-gray-800 overflow-hidden">
                    <CommandInput
                      placeholder="Digite um comando..."
                      value={commandInput}
                      onValueChange={setCommandInput}
                      onKeyDown={(e: React.KeyboardEvent) => {
                        if (e.key === "Enter") {
                          executeCommand();
                        }
                      }}
                    />
                  </Command>
                </div>
                <Button size="icon" onClick={executeCommand}>
                  <ArrowRight className="h-4 w-4" />
                </Button>
              </div>
            </TabsContent>

            <TabsContent value="commands" className="flex-1 p-4 m-0">
              <div className="grid grid-cols-2 gap-4">
                {commonCommands.map((cmd) => (
                  <div
                    key={cmd.name}
                    className="bg-gray-900 p-4 rounded-lg border border-gray-800 hover:border-indigo-500 cursor-pointer"
                    onClick={() => selectCommand(cmd.name)}
                  >
                    <div className="flex items-center justify-between">
                      <code className="text-indigo-400 font-mono">
                        {cmd.name}
                      </code>
                      <Play className="h-4 w-4 text-gray-400" />
                    </div>
                    <p className="text-sm text-gray-400 mt-1">
                      {cmd.description}
                    </p>
                  </div>
                ))}
              </div>
            </TabsContent>
          </Tabs>
        </div>
      </div>

      {/* Barra de status */}
      <div className="bg-gray-900 border-t border-gray-800 p-2 text-sm text-gray-500 flex justify-between">
        <span>Conectado: Local</span>
        <span>Terminal GUI v1.0.0</span>
      </div>
    </div>
  );
}
