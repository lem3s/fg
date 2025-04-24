interface Window {
    go: {
        main: {
            App: {
                ExecuteCommand: (command: string) => Promise<{
                    output: string;
                    currentDirectory?: string;
                    error?: string;
                }>;
                GetCurrentDirectory: () => Promise<string>;
                OpenTerminalHere: (path?: string) => Promise<{
                    output: string;
                    currentDirectory?: string;
                    error?: string;
                }>;
            }
        }
    }
}