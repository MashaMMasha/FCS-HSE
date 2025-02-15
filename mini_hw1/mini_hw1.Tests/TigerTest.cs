using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using mini_hw1;

namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Tiger))]
public class TigerTest
{

    [TestMethod]
    public void TestIsHealthyTrue_Young()
    {
        var tiger = new Tiger("Tiger",  5, 3,8, 1, 20);
        var result = tiger.IsHealthy();
        Assert.IsTrue(result);
        Assert.IsTrue(tiger.Name == "Tiger");
    }
    [TestMethod]
    public void TestIsHealthyTrue_HasMeat()
    {
        var tiger = new Tiger("Tiger",  5, 20,8, 1, 20, 10);
        var result = tiger.IsHealthy();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestIsHealthyFalse()
    {
        var tiger = new Tiger("Tiger",  5, 20, 8, 1, 20);
        var result = tiger.IsHealthy();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_true()
    {
        var tiger = new Tiger("Tiger",  5, 3, 1, 19, 20, 5);
        var result = tiger.CanBeInContactZoo();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_False()
    {
        var tiger = new Tiger("Tiger",  5, 3, 10, 19, 20, 0);
        var result = tiger.CanBeInContactZoo();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestGiveMeat()
    {
        var tiger = new Tiger("Tiger",  5, 3, 10, 19, 20, 0);
        int meat = tiger.Meat;
        tiger.GiveFood();
        Assert.AreEqual(meat + 1, tiger.Meat);
    }
}