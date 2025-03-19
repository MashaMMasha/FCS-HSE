using System;
using System.Collections.Generic;
using System.Linq;
namespace BankSystem;

public class BankFacade
{
    private Dictionary<Guid, BankAccount> _accounts;
    private Dictionary<Guid, Category> _categories;
    private Dictionary<Guid, Operation> _operations;

    public BankFacade()
    {
        _accounts = new Dictionary<Guid, BankAccount>();
        _categories = new Dictionary<Guid, Category>();
        _operations = new Dictionary<Guid, Operation>();
    }

    public Guid CreateAccount(string name)
    {
        var account = new BankAccount(name);
        _accounts.Add(account.id, account);
        return account.id;
    }

    public Guid CreateAccount(string name, double balance)
    {
        var account = new BankAccount(name, balance);
        _accounts.Add(account.id, account);
        return account.id;
    }

    public BankAccount GetAccount(Guid id)
    {
        return _accounts[id];
    }

    public void DeleteAccount(Guid id)
    {
        _accounts.Remove(id);
    }

    public Guid CreateCategory(string name, OperationType type)
    {
        var category = new Category(name, type);
        _categories.Add(category.id, category);
        return category.id;
    }

    public void RenameCategory(Guid id, string newName)
    {
        if (_categories.ContainsKey(id))
        {
            _categories[id].name = newName;
        }
    }

    public void DeleteCategory(Guid id)
    {
        _categories.Remove(id);
    }

    public Guid AddOperation(Guid accountId, Guid categoryId, OperationType type, int amount, DateTime date,
        string description = "")
    {
        var operation = new Operation(accountId, categoryId, type, amount, date, description);
        _operations.Add(operation.id, operation);
        return operation.id;
    }

    public void ChangeOperationDescription(Guid id, string newDescription)
    {
        if (_operations.ContainsKey(id))
        {
            _operations[id].description = newDescription;
        }
    }

    public void ChangeOperationCategory(Guid id, Guid newCategoryId)
    {
        if (_operations.ContainsKey(id))
        {
            _operations[id].categoryId = newCategoryId;
        }
    }

    public void DeleteOperation(Guid id)
    {
        _operations.Remove(id);
    }

    public IEnumerable<Operation> GetOperationsByAccount(Guid accountId)
    {
        return _operations.Values.Where(op => op.accountId == accountId);
    }

    public IEnumerable<Operation> GetOperationsByCategory(Guid categoryId, Guid accountId)
    {
        return _operations.Values.Where(op => op.categoryId == categoryId && op.accountId == accountId);
    }

    public IEnumerable<Operation> GetOperationsByDateRange(DateTime start, DateTime end)
    {
        return _operations.Values.Where(op => op.date >= start && op.date <= end);
    }

    public IEnumerable<Operation> GetOperationsByDateAndCategory(DateTime start, DateTime end, Guid categoryId)
    {
        return _operations.Values.Where(op => op.date >= start && op.date <= end && op.categoryId == categoryId);
    }
}