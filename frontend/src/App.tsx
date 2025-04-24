import { useState, useCallback, useEffect, useRef } from "react";
import { 
  Terminal, 
  Code, 
  Play 
} from "lucide-react";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

export default function TerminalGUI() {
  const [currentDirectory, setCurrentDirectory] = useState('');
  const [_, setCommandHistory] = useState<string[]>([]);
  const [output, setOutput] = useState<string[]>([]);
  const [commandInput, setCommandInput] = useState('');
  const inputRef = useRef<HTMLInputElement>(null);
  const [activeTab, setActiveTab] = useState('terminal');

  const customCommands = [
    {
      category: "Environment",
      commands: [
        { 
          name: "ls", 
          description: "Listar ambientes configurados",
          example: "ls"
        },
      ]
    },
    
  ];

  const executeCommand = useCallback(async (cmd?: string) => {
    const commandToExecute = cmd || commandInput;
    
    if (!commandToExecute.trim()) return;

    if (commandToExecute.trim().startsWith('terminal ')) {
      try {
        const path = commandToExecute.trim().split(' ')[1] || '.';
        const result = await window.go.main.App.OpenTerminalHere(path);

        setCommandHistory(prev => [...prev, commandToExecute]);
        setOutput(prev => [
          ...prev, 
          `> ${commandToExecute}`,
          result.output,
          result.error || ''
        ]);

        if (result.currentDirectory) {
          setCurrentDirectory(result.currentDirectory);
        }

        setCommandInput('');
        return;
      } catch (error) {
        console.error('Erro ao abrir terminal:', error);
      }
    }

    if (commandToExecute.trim() === 'clear') {
      setOutput([]);
      setCommandInput('');
      return;
    }

    try {
      const result = await window.go.main.App.ExecuteCommand(commandToExecute);

      setCommandHistory(prev => [...prev, commandToExecute]);
      setOutput(prev => [
        ...prev, 
        `> ${commandToExecute}`,
        result.output,
        result.error || ''
      ]);

      if (result.currentDirectory) {
        setCurrentDirectory(result.currentDirectory);
      }

      setCommandInput('');
      
      setActiveTab('terminal');
    } catch (error) {
      console.error('Erro ao executar comando:', error);
    }
  }, [commandInput]);

  useEffect(() => {
    inputRef.current?.focus();
  }, [output]);

  useEffect(() => {
    const fetchCurrentDirectory = async () => {
      try {
        const dir = await window.go.main.App.GetCurrentDirectory();
        setCurrentDirectory(dir);
      } catch (error) {
        console.error('Erro ao buscar diretório:', error);
      }
    };

    fetchCurrentDirectory();
  }, []);

  const handleCommandSelect = useCallback((cmd: string) => {
    setCommandInput(cmd);
    setActiveTab('terminal');
    
    executeCommand(cmd);
  }, [executeCommand]);

  return (
    <div className="min-h-screen bg-gray-950 text-gray-200 flex flex-col">
      <div className="flex-1 flex flex-col">
        <Tabs 
          value={activeTab} 
          onValueChange={setActiveTab} 
          defaultValue="terminal" 
          className="flex-1 flex flex-col"
        >
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
            <div className="flex-1 bg-gray-950 overflow-auto p-4 font-mono text-sm">
              {output.map((line, index) => (
                <div key={index} className="whitespace-pre-wrap">
                  {line}
                </div>
              ))}
              <div className="flex items-center">
                <span className="mr-2 text-green-400">{currentDirectory}</span>
                <input
                  ref={inputRef}
                  type="text"
                  value={commandInput}
                  onChange={(e) => setCommandInput(e.target.value)}
                  onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                      executeCommand();
                    }
                  }}
                  className="flex-1 bg-transparent outline-none text-white"
                  placeholder="Digite um comando..."
                />
              </div>
            </div>
          </TabsContent>

          <TabsContent value="commands" className="flex-1 p-4 m-0">
            <div className="space-y-6">
              {customCommands.map((category, catIndex) => (
                <div key={catIndex}>
                  <h2 className="text-lg font-semibold mb-3 text-gray-300">
                    {category.category}
                  </h2>
                  <div className="grid grid-cols-2 gap-4">
                    {category.commands.map((cmd, cmdIndex) => (
                      <div
                        key={cmdIndex}
                        className="bg-gray-900 p-4 rounded-lg border border-gray-800 hover:border-indigo-500 cursor-pointer transition-all duration-300"
                        onClick={() => handleCommandSelect(cmd.example)}
                      >
                        <div className="flex items-center justify-between mb-2">
                          <code className="text-indigo-400 font-mono">
                            {cmd.name}
                          </code>
                          <Play className="h-4 w-4 text-gray-400" />
                        </div>
                        <p className="text-sm text-gray-400">
                          {cmd.description}
                        </p>
                        <div className="mt-2 text-xs text-gray-600">
                          Exemplo: <code>{cmd.example}</code>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
}
