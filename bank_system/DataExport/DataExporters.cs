using System;
using System.Collections.Generic;

namespace BankSystem;

public class JsonDataExporter : DataExporter
{
    protected override string SerializeData(List<BankAccount> accounts, List<Category> categories, List<Operation> operations)
    {
        var accountsJson = string.Join(",", accounts.ConvertAll(a => $"{{\"name\": \"{a.name}\", \"balance\": {a.balance}}}"));
        var categoriesJson = string.Join(",", categories.ConvertAll(c => $"{{\"name\": \"{c.name}\", \"type\": \"{c.type}\"}}"));
        var operationsJson = string.Join(",", operations.ConvertAll(o => $"{{\"description\": \"{o.description}\", \"amount\": {o.amount}}}"));

        return $"{{\"accounts\": [{accountsJson}], \"categories\": [{categoriesJson}], \"operations\": [{operationsJson}]}}";
    }
}

public class YamlDataExporter : DataExporter
{
    protected override string SerializeData(List<BankAccount> accounts, List<Category> categories, List<Operation> operations)
    {
        var accountsYaml = string.Join("\n", accounts.ConvertAll(a => $"- name: {a.name}\n  balance: {a.balance}"));
        var categoriesYaml = string.Join("\n", categories.ConvertAll(c => $"- name: {c.name}\n  type: {c.type}"));
        var operationsYaml = string.Join("\n", operations.ConvertAll(o => $"- description: {o.description}\n  amount: {o.amount}"));

        return $"accounts:\n{accountsYaml}\ncategories:\n{categoriesYaml}\noperations:\n{operationsYaml}";
    }
}

public class CsvDataExporter : DataExporter
{
    protected override string SerializeData(List<BankAccount> accounts, List<Category> categories,
        List<Operation> operations)
    {
        var accountsCsv = string.Join("\n", accounts.ConvertAll(a => $"BankAccount,{a.name},{a.balance}"));
        var categoriesCsv = string.Join("\n", categories.ConvertAll(c => $"Category,{c.name},{c.type}"));
        var operationsCsv = string.Join("\n", operations.ConvertAll(o => $"Operation,{o.description},{o.amount}"));

        return $"Type,Name,Value\n{accountsCsv}\n{categoriesCsv}\n{operationsCsv}";
    }
}