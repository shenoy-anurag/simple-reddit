describe('Navigation Menu Tests', () => {
    it('Open Home Page', () => {
        cy.visit('http://localhost:4200/');

        cy.contains('Home').click();
        cy.url().should('include', '/home');
    })

    it('Open Subreddits', () => {
        cy.contains('Subreddits').click();
        cy.url().should('include', '/subreddits');
    })

    it('Reddit Logo Home Button', () => {
        cy.contains('Reddit').click();
        cy.url().should('include', '/home');
    })
})

describe('Sign Up User', () => {
    it('Navigate to Log In', () => {
        cy.contains('Log In').click();
        cy.url().should('include', '/login');
    })

    it('Navigate to Sign Up', () => {
        cy.contains('Sign Up').click();
        cy.url().should('include', '/signup');
    })

    it('Populate Form', () => {
        cy.get('input[name="firstname"]').type('John').should('have.value', "John");
        cy.get('input[name="lastname"]').type('Doe').should('have.value', "Doe");
        cy.get('input[name="email"]').type('John.Doe@email.com').should('have.value', "John.Doe@email.com");
        cy.get('input[name="username"]').type('JohnDoe123').should('have.value', "JohnDoe123");
        cy.get('input[name="password"]').type('Aa1abc').should('have.value', "Aa1abc");
        cy.get('input[name="password2"]').type('Aa1abc').should('have.value', "Aa1abc");

        cy.contains('Submit').click();
        cy.url().should('include', '/popup-message');

        cy.get('h1').should('contain', 'Succesfully Signed-Up');
        cy.contains('Continue').click();
    })
})
describe('Sign In User', () => {
    it('Navigate to Log In', () => {
        cy.contains('Log In').click();
        cy.url().should('include', '/login')
    })

    it('Populate Form', () => {
        cy.get('input[name="username"]').type("JohnDoe123").should('have.value', 'JohnDoe123');
        cy.get('input[name="password"]').type('Aa1abc').should('have.value', "Aa1abc");
    })

    it('Log In', () => {
        cy.get('button[name="loginbutton"]').click();
    })

    it('Check Profile', () => {
        cy.contains('Profile').click();
        cy.url().should('include', '/profile')
    })

    it('Log Out', () => {
        cy.contains("Log Out").click();
        cy.url().should('include', '/login')
    })
})

describe('Create Post', () => {
    it('Navigate to Home', () => {
        cy.contains('Home').click();
        cy.url().should('include', '/home');
    })

    it('Create New Post', () => {
        cy.contains('Create New Post').click();
        cy.url().should('include', '/newpostform');
    })

    it('Populate Form', () => {
        cy.get('input[name="username"]').type("JohnDoe123").should('have.value', 'JohnDoe123');
        cy.get('input[name="title"]').type('Cypress Community Title').should('have.value', "Cypress Community Title");
    })
})
