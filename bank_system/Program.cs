using System;
using Microsoft.Extensions.DependencyInjection;
namespace BankSystem;

class Program
{
    static IServiceProvider ConfigureServices()
    {
        var services = new ServiceCollection();
        
        services.AddSingleton<BankAccountFactory>();
        services.AddSingleton<CategoryFactory>();
        services.AddSingleton<OperationFactory>();
        services.AddSingleton<BankFacade>();

        return services.BuildServiceProvider();
    }
    static void Main()
    {
        var serviceProvider = ConfigureServices();
        var bank = serviceProvider.GetRequiredService<BankFacade>();

        Guid accountId = bank.CreateAccount("Мурашко Мария основной счет", 5000);
        Guid accountId2 = bank.CreateAccount("Мурашко Мария дополнительный счет", 5000);

        Guid categoryId = bank.CreateCategory("Продукты", OperationType.Income);

        Guid operationId = bank.AddOperation(accountId, categoryId, OperationType.Income, 200, DateTime.Now, "Зарплата");
        bool deleted = bank.DeleteAccount(Guid.NewGuid()); // ожидается ошибка так как такого счета не существует
        
        var accounts = bank.GetAllAccounts();
        var categories = bank.GetAllCategories();
        var operations = bank.GetAllOperations();
        
        var jsonExporter = new JsonDataExporter();
        
        var consoleObserver = new ConsoleObserver();
        jsonExporter.AddObserver(consoleObserver);
        
        jsonExporter.Export("export.json", accounts, categories, operations);
    }
}
