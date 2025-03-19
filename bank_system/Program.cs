using BankSystem;
using Microsoft.Extensions.DependencyInjection;

class Program
{
    static void Main()
    {
        var serviceProvider = new ServiceCollection().AddSingleton<BankFacade>().BuildServiceProvider();
        var bank = serviceProvider.GetRequiredService<BankFacade>();
        var commandQueue = new CommandQueue();

        
        commandQueue.AddCommand(new BankCommand(() => bank.CreateAccount("1000"), "Создание счета"));
        
        commandQueue.ProcessCommands();
    }
}