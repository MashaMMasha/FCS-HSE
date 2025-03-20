using System;
using System.Collections.Generic;

namespace BankSystem;

public class BankFacade
{
    private readonly BankAccountFactory _accountFactory;
    private readonly CategoryFactory _categoryFactory;
    private readonly OperationFactory _operationFactory;

    public BankFacade(
        BankAccountFactory accountFactory,
        CategoryFactory categoryFactory,
        OperationFactory operationFactory)
    {
        _accountFactory = accountFactory;
        _categoryFactory = categoryFactory;
        _operationFactory = operationFactory;
    }

    public Guid CreateAccount(string name)
    {
        var command = new BankCommand<Guid>(() => _accountFactory.CreateAccount(name));
        var loggedCommand = new LoggingCommand<Guid>(command, $"Создание счета: {name}");
        return loggedCommand.Execute();
    }

    public Guid CreateAccount(string name, double balance)
    {
        var command = new BankCommand<Guid>(() => _accountFactory.CreateAccount(name, balance));
        var loggedCommand = new LoggingCommand<Guid>(command, $"Создание счета: {name} с балансом {balance}");
        return loggedCommand.Execute();
    }

    public BankAccount GetAccount(Guid id)
    {
        return _accountFactory.GetAccount(id);
    }

    public bool DeleteAccount(Guid id)
    {
        var command = new BankCommand<bool>(() => _accountFactory.DeleteAccount(id));
        var loggedCommand = new LoggingCommand<bool>(command, $"Удаление счета: {id}");
        return loggedCommand.Execute();
    }

    public Guid CreateCategory(string name, OperationType type)
    {
        var command = new BankCommand<Guid>(() => _categoryFactory.CreateCategory(name, type));
        var loggedCommand = new LoggingCommand<Guid>(command, $"Создание категории: {name} ({type})");
        return loggedCommand.Execute();
    }

    public bool RenameCategory(Guid id, string newName)
    {
        var command = new BankCommand<bool>(() => _categoryFactory.RenameCategory(id, newName));
        var loggedCommand = new LoggingCommand<bool>(command, $"Переименование категории: {id} -> {newName}");
        return loggedCommand.Execute();
    }

    public bool DeleteCategory(Guid id)
    {
        var command = new BankCommand<bool>(() => _categoryFactory.DeleteCategory(id));
        var loggedCommand = new LoggingCommand<bool>(command, $"Удаление категории: {id}");
        return loggedCommand.Execute();
    }

    public Guid AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date,
        string description = "")
    {
        var command = new BankCommand<Guid>(
            () => _operationFactory.AddOperation(accountId, categoryId, type, amount, date, description));
        var loggedCommand = new LoggingCommand<Guid>(command, $"Добавление операции: {type} на сумму {amount} ({description})");
        return loggedCommand.Execute();
    }

    public IEnumerable<Operation> GetOperationsByAccount(Guid accountId)
    {
        return _operationFactory.GetOperations(accountId);
    }

    public IEnumerable<Operation> GetOperationsByCategory(Guid categoryId, Guid accountId)
    {
        return _operationFactory.GetOperations(categoryId, accountId);
    }

    public IEnumerable<Operation> GetOperationsByDateRange(DateTime start, DateTime end)
    {
        return _operationFactory.GetOperations(start, end);
    }

    public IEnumerable<Operation> GetOperationsByDateAndCategory(DateTime start, DateTime end, Guid categoryId)
    {
        return _operationFactory.GetOperations(start, end, categoryId);
    }
    public List<Category> GetAllCategories()
    {
        return _categoryFactory.GetAllCategories();
    }
    public List<BankAccount> GetAllAccounts()
    {
        return _accountFactory.GetAllAccounts();
    }
    public List<Operation> GetAllOperations()
    {
        return _operationFactory.GetOperations();
    }
}
