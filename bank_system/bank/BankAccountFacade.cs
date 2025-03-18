using System.Collections.Generic;

namespace BankSystem;
public class BankAccountFacade
{
    private Dictionary<Guid, BankAccount> accounts = new();

    public Guid CreateAccount(string name)
    {
        var account = new BankAccount(name);
        accounts.Add(account.id, account);
    }
    public Guid CreateAccount(string name, double balance)
    {
        var account = new BankAccount(name, balance);
        accounts.Add(account.id, account);
    }
    public BankAccount GetAccount(Guid id)
    {
        return accounts[id];
    }
    public void DeleteAccount(Guid id)
    {
        accounts.Remove(id);
    }
}