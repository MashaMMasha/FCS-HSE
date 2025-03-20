namespace BankSystem;

using System.Collections.Generic;
public class BankAccountFactory
{
    private Dictionary<Guid, BankAccount> _accounts = new();

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
    public bool DeleteAccount(Guid id)
    {
        if (!_accounts.ContainsKey(id))
        {
            return false;
        }
        _accounts.Remove(id);
        return true;
    }
    public List<BankAccount> GetAllAccounts()
    {
        return new List<BankAccount>(_accounts.Values);
    }
}